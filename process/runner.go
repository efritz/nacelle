package process

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/efritz/nacelle/config"
	"github.com/efritz/nacelle/logging"
	"github.com/efritz/nacelle/service"
)

type (
	Container interface {
		RegisterInitializer(Initializer, ...InitializerConfigFunc)
		RegisterProcess(Process, ...ProcessConfigFunc)
	}

	Runner interface {
		Container
		Run(config.Config, logging.Logger) <-chan error
	}

	// ProcessRunner maintains a set of registered initializers and processes,
	// starts them in order, and then monitors their results.
	runner struct {
		container       service.Container
		initializers    []*initializerMeta
		processes       map[int][]*processMeta
		done            chan struct{}
		halt            chan struct{}
		once            *sync.Once
		shutdownTimeout time.Duration
	}

	errMeta struct {
		err     error
		process *processMeta
	}
)

var ErrInitTimeout = fmt.Errorf("init method did not finish within timeout")

// NewRunner creates a new process runner with the given service container.
func NewRunner(
	container service.Container,
	runnerConfigs ...RunnerConfigFunc,
) Runner {
	runner := &runner{
		container:    container,
		initializers: []*initializerMeta{},
		processes:    map[int][]*processMeta{},
		done:         make(chan struct{}),
		halt:         make(chan struct{}),
		once:         &sync.Once{},
	}

	for _, f := range runnerConfigs {
		f(runner)
	}

	return runner
}

// RegisterInitializer registers an initializer with the given configuration. The
// order the initializers are run mirrors the order of registration.
func (pr *runner) RegisterInitializer(
	initializer Initializer,
	initializerConfigs ...InitializerConfigFunc,
) {
	meta := newInitializerMeta(initializer)

	for _, f := range initializerConfigs {
		f(meta)
	}

	pr.initializers = append(pr.initializers, meta)
}

// RegisterProcess registers a process with the given configuration. The order
// of process registration is arbitrary.
func (pr *runner) RegisterProcess(
	process Process,
	processConfigs ...ProcessConfigFunc,
) {
	meta := newProcessMeta(process)

	for _, f := range processConfigs {
		f(meta)
	}

	if _, ok := pr.processes[meta.priority]; !ok {
		pr.processes[meta.priority] = []*processMeta{}
	}

	pr.processes[meta.priority] = append(pr.processes[meta.priority], meta)
}

// Run will run the registered initializers and processes with the given loaded
// configuration object. It will return a read-only channel of error values on
// which non-nil error results from initializers and proceses are written.
//
// For each initializer, in order of registration: services are injected into the
// initializer and then its Init method is called. Initializers are run one at a
// time and an error from an initializer will cause an immediate return from Run.
//
// For each processes set with the same priority (lowest to highest): services are
// injected into each process and each Init method is called. Init methods are called
// one at a time and in the order of process registration. If an Init method returns
// an error, all lower-priority processes are stopped. Then, the Start method for each
// process is called concurrently in its own goroutine.
//
// If any process returns a non-nil error from Start, all running processes will be
// stopped. If a process return a nil error and has not been configured for silent exit,
// the same behavior will occur.
//
// Receiving an external signal (SIGINT or SIGTERM) will also start a graceful shutdown.
// A second signal will cause the Run method to stop blocking (although a process may
// still be running in a goroutine).
//
// If any process has started, the error channel returned from Run will remain open
// until all running processes have exited.
func (pr *runner) Run(config config.Config, logger logging.Logger) <-chan error {
	n := 0
	for _, ps := range pr.processes {
		n += len(ps)
	}

	errChan := make(chan error, n*2+1)

	if err := pr.runInitializers(config, logger); err != nil {
		defer close(errChan)
		errChan <- err
		return errChan
	}

	var (
		startErrors = make(chan errMeta)
		priorities  = pr.getPriorities()
		wg          = &sync.WaitGroup{}
	)

	if !pr.runProcesses(
		priorities,
		config,
		logger,
		startErrors,
		errChan,
		wg,
	) {
		return errChan
	}

	logger.Info("All processes running")
	go pr.watch(priorities, logger, startErrors, errChan)
	go closeAfterWait(wg, startErrors)
	return chainUntilHalt(errChan, pr.done)
}

// Shutdown initiates a graceful shutdown of processes (if it is not already
// within a graceful shutdown period). This method will block until the runner
// has exited or the timeout period has elapsed. In the later case, an error
// is returned.
func (pr *runner) Shutdown(timeout time.Duration) error {
	pr.once.Do(func() {
		close(pr.halt)
	})

	select {
	case <-time.After(timeout):
		return fmt.Errorf("process runner did not complete within timeout")
	case <-pr.done:
		return nil
	}
}

//
// Internals

func (pr *runner) getPriorities() []int {
	priorities := []int{}
	for priority := range pr.processes {
		priorities = append(priorities, priority)
	}

	sort.Ints(priorities)
	return priorities
}

func (pr *runner) runInitializers(config config.Config, logger logging.Logger) error {
	logger.Info("Running initializers")

	for _, initializer := range pr.initializers {
		logger.Debug("Injecting services into %s", initializer.Name())

		if err := pr.container.Inject(initializer.Initializer); err != nil {
			return fmt.Errorf(
				"failed to inject services into %s (%s)",
				initializer.Name(),
				err.Error(),
			)
		}

		logger.Debug("Initializing %s", initializer.Name())

		if err := initWithTimeout(initializer, config, initializer.timeout); err != nil {
			return fmt.Errorf(
				"failed to initialize %s (%s)",
				initializer.Name(),
				err.Error(),
			)
		}

		logger.Debug("Initialized %s", initializer.Name())
	}

	return nil
}

func (pr *runner) runProcesses(
	priorities []int,
	config config.Config,
	logger logging.Logger,
	startErrors chan errMeta,
	errChan chan error,
	wg *sync.WaitGroup,
) bool {
	logger.Debug("Injecting services into process instances")

	for i := range priorities {
		for _, process := range pr.processes[priorities[i]] {
			if err := pr.container.Inject(process.Process); err != nil {
				defer close(errChan)

				errChan <- fmt.Errorf(
					"failed to inject services into %s (%s)",
					process.Name(),
					err.Error(),
				)

				return false
			}
		}
	}

	logger.Info("Initializing and starting processes")

	for i := range priorities {
		err := pr.initAndStartProcesses(
			pr.processes[priorities[i]],
			priorities[i],
			config,
			logger,
			wg,
			startErrors,
		)

		if err != nil {
			errChan <- err
			pr.stopProcesessBelowPriority(priorities, i, logger, errChan)
			go closeAfterWait(wg, startErrors)

			go func() {
				defer close(errChan)

				for err := range startErrors {
					if err.err != nil {
						errChan <- err.err
					}
				}
			}()

			return false
		}
	}

	return true
}

func (pr *runner) initAndStartProcesses(
	processes []*processMeta,
	priority int,
	config config.Config,
	logger logging.Logger,
	wg *sync.WaitGroup,
	startErrors chan<- errMeta,
) error {
	logger.Debug("Initializing processes at priority %d", priority)

	for _, process := range processes {
		logger.Debug("Initializing %s", process.Name())

		if err := initWithTimeout(process, config, process.initTimeout); err != nil {
			return fmt.Errorf("failed to initialize %s (%s)", process.Name(), err.Error())
		}

		logger.Debug("Initialized %s", process.Name())
	}

	logger.Debug("Starting processes at priority %d", priority)

	for _, process := range processes {
		wg.Add(1)

		go func(process *processMeta) {
			defer wg.Done()

			logger.Debug("Starting %s", process.Name())

			err := process.Start()
			if err != nil {
				err = fmt.Errorf("%s returned a fatal error (%s)", process.Name(), err.Error())
			}

			startErrors <- errMeta{err, process}
		}(process)
	}

	return nil
}

func (pr *runner) watch(
	priorities []int,
	logger logging.Logger,
	startErrors <-chan errMeta,
	errChan chan<- error,
) {
	defer close(pr.done)

	callback := func() {
		pr.stopProcesessBelowPriority(
			priorities,
			len(priorities),
			logger,
			errChan,
		)
	}

	watcher := newWatcher(
		callback,
		logger,
		pr.shutdownTimeout,
		startErrors,
		errChan,
		pr.halt,
	)

	watcher.watch()
}

func (pr *runner) stopProcesessBelowPriority(priorities []int, p int, logger logging.Logger, errChan chan<- error) {
	for i := p - 1; i >= 0; i-- {
		pr.stopProcesses(pr.processes[priorities[i]], priorities[i], logger, errChan)
	}
}

func (pr *runner) stopProcesses(processes []*processMeta, priority int, logger logging.Logger, errChan chan<- error) {
	logger.Debug("Stopping processes at priority %d", priority)

	for _, process := range processes {
		logger.Debug("Stopping %s", process.Name())

		process.once.Do(func() {
			if err := process.Stop(); err != nil {
				errChan <- fmt.Errorf("%s returned error from stop (%s)", process.Name(), err.Error())
			}
		})
	}
}

//
// Helpers

func initWithTimeout(initializer Initializer, config config.Config, timeout time.Duration) error {
	ch := make(chan error)

	go func() {
		defer close(ch)
		ch <- initializer.Init(config)
	}()

	select {
	case err := <-ch:
		return err
	case <-makeTimeoutChan(timeout):
		return ErrInitTimeout
	}
}

func makeTimeoutChan(timeout time.Duration) <-chan time.Time {
	if timeout == 0 {
		return nil
	}

	return time.After(timeout)
}

func closeAfterWait(wg *sync.WaitGroup, startErrors chan errMeta) {
	wg.Wait()
	close(startErrors)
}

func chainUntilHalt(src <-chan error, halt <-chan struct{}) <-chan error {
	out := make(chan error)

	go func() {
	loop:
		for {
			select {
			case err, ok := <-src:
				if !ok {
					break loop
				}

				out <- err

			case <-halt:
				break loop
			}
		}

		close(out)

		for range src {
		}
	}()

	return out
}