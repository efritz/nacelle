// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-16T11:29:10-05:00
// using the command
// $ go-mockgen github.com/efritz/nacelle/config -i Config -o mock_config_test.go -f

package http

import (
	config "github.com/efritz/nacelle/config"
	tags "github.com/efritz/zubrin/tags"
)

// MockConfig is a mock impelementation of the Config interface (from the
// package github.com/efritz/nacelle/config) used for unit testing.
type MockConfig struct {
	// AssetsFunc is an instance of a mock function object controlling the
	// behavior of the method Assets.
	AssetsFunc *ConfigAssetsFunc
	// DumpFunc is an instance of a mock function object controlling the
	// behavior of the method Dump.
	DumpFunc *ConfigDumpFunc
	// LoadFunc is an instance of a mock function object controlling the
	// behavior of the method Load.
	LoadFunc *ConfigLoadFunc
	// MustLoadFunc is an instance of a mock function object controlling the
	// behavior of the method MustLoad.
	MustLoadFunc *ConfigMustLoadFunc
}

// NewMockConfig creates a new mock of the Config interface. All methods
// return zero values for all results, unless overwritten.
func NewMockConfig() *MockConfig {
	return &MockConfig{
		AssetsFunc: &ConfigAssetsFunc{
			defaultHook: func() []string {
				return nil
			},
		},
		DumpFunc: &ConfigDumpFunc{
			defaultHook: func() map[string]string {
				return nil
			},
		},
		LoadFunc: &ConfigLoadFunc{
			defaultHook: func(interface{}, ...tags.TagModifier) error {
				return nil
			},
		},
		MustLoadFunc: &ConfigMustLoadFunc{
			defaultHook: func(interface{}, ...tags.TagModifier) {
				return
			},
		},
	}
}

// NewMockConfigFrom creates a new mock of the MockConfig interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockConfigFrom(i config.Config) *MockConfig {
	return &MockConfig{
		AssetsFunc: &ConfigAssetsFunc{
			defaultHook: i.Assets,
		},
		DumpFunc: &ConfigDumpFunc{
			defaultHook: i.Dump,
		},
		LoadFunc: &ConfigLoadFunc{
			defaultHook: i.Load,
		},
		MustLoadFunc: &ConfigMustLoadFunc{
			defaultHook: i.MustLoad,
		},
	}
}

// ConfigAssetsFunc describes the behavior when the Assets method of the
// parent MockConfig instance is invoked.
type ConfigAssetsFunc struct {
	defaultHook func() []string
	hooks       []func() []string
	history     []ConfigAssetsFuncCall
}

// Assets delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) Assets() []string {
	r0 := m.AssetsFunc.nextHook()()
	m.AssetsFunc.history = append(m.AssetsFunc.history, ConfigAssetsFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Assets method of the
// parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigAssetsFunc) SetDefaultHook(hook func() []string) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Assets method of the parent MockConfig instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigAssetsFunc) PushHook(hook func() []string) {
	f.hooks = append(f.hooks, hook)
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigAssetsFunc) SetDefaultReturn(r0 []string) {
	f.SetDefaultHook(func() []string {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigAssetsFunc) PushReturn(r0 []string) {
	f.PushHook(func() []string {
		return r0
	})
}

func (f *ConfigAssetsFunc) nextHook() func() []string {
	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

// History returns a sequence of ConfigAssetsFuncCall objects describing the
// invocations of this function.
func (f *ConfigAssetsFunc) History() []ConfigAssetsFuncCall {
	return f.history
}

// ConfigAssetsFuncCall is an object that describes an invocation of method
// Assets on an instance of MockConfig.
type ConfigAssetsFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ConfigAssetsFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigAssetsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ConfigDumpFunc describes the behavior when the Dump method of the parent
// MockConfig instance is invoked.
type ConfigDumpFunc struct {
	defaultHook func() map[string]string
	hooks       []func() map[string]string
	history     []ConfigDumpFuncCall
}

// Dump delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) Dump() map[string]string {
	r0 := m.DumpFunc.nextHook()()
	m.DumpFunc.history = append(m.DumpFunc.history, ConfigDumpFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Dump method of the
// parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigDumpFunc) SetDefaultHook(hook func() map[string]string) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Dump method of the parent MockConfig instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigDumpFunc) PushHook(hook func() map[string]string) {
	f.hooks = append(f.hooks, hook)
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigDumpFunc) SetDefaultReturn(r0 map[string]string) {
	f.SetDefaultHook(func() map[string]string {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigDumpFunc) PushReturn(r0 map[string]string) {
	f.PushHook(func() map[string]string {
		return r0
	})
}

func (f *ConfigDumpFunc) nextHook() func() map[string]string {
	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

// History returns a sequence of ConfigDumpFuncCall objects describing the
// invocations of this function.
func (f *ConfigDumpFunc) History() []ConfigDumpFuncCall {
	return f.history
}

// ConfigDumpFuncCall is an object that describes an invocation of method
// Dump on an instance of MockConfig.
type ConfigDumpFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 map[string]string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ConfigDumpFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigDumpFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ConfigLoadFunc describes the behavior when the Load method of the parent
// MockConfig instance is invoked.
type ConfigLoadFunc struct {
	defaultHook func(interface{}, ...tags.TagModifier) error
	hooks       []func(interface{}, ...tags.TagModifier) error
	history     []ConfigLoadFuncCall
}

// Load delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) Load(v0 interface{}, v1 ...tags.TagModifier) error {
	r0 := m.LoadFunc.nextHook()(v0, v1...)
	m.LoadFunc.history = append(m.LoadFunc.history, ConfigLoadFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Load method of the
// parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigLoadFunc) SetDefaultHook(hook func(interface{}, ...tags.TagModifier) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Load method of the parent MockConfig instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigLoadFunc) PushHook(hook func(interface{}, ...tags.TagModifier) error) {
	f.hooks = append(f.hooks, hook)
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigLoadFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(interface{}, ...tags.TagModifier) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigLoadFunc) PushReturn(r0 error) {
	f.PushHook(func(interface{}, ...tags.TagModifier) error {
		return r0
	})
}

func (f *ConfigLoadFunc) nextHook() func(interface{}, ...tags.TagModifier) error {
	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

// History returns a sequence of ConfigLoadFuncCall objects describing the
// invocations of this function.
func (f *ConfigLoadFunc) History() []ConfigLoadFuncCall {
	return f.history
}

// ConfigLoadFuncCall is an object that describes an invocation of method
// Load on an instance of MockConfig.
type ConfigLoadFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 interface{}
	// Arg1 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg1 []tags.TagModifier
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c ConfigLoadFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg1 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigLoadFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ConfigMustLoadFunc describes the behavior when the MustLoad method of the
// parent MockConfig instance is invoked.
type ConfigMustLoadFunc struct {
	defaultHook func(interface{}, ...tags.TagModifier)
	hooks       []func(interface{}, ...tags.TagModifier)
	history     []ConfigMustLoadFuncCall
}

// MustLoad delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) MustLoad(v0 interface{}, v1 ...tags.TagModifier) {
	m.MustLoadFunc.nextHook()(v0, v1...)
	m.MustLoadFunc.history = append(m.MustLoadFunc.history, ConfigMustLoadFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the MustLoad method of
// the parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigMustLoadFunc) SetDefaultHook(hook func(interface{}, ...tags.TagModifier)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// MustLoad method of the parent MockConfig instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigMustLoadFunc) PushHook(hook func(interface{}, ...tags.TagModifier)) {
	f.hooks = append(f.hooks, hook)
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigMustLoadFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(interface{}, ...tags.TagModifier) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigMustLoadFunc) PushReturn() {
	f.PushHook(func(interface{}, ...tags.TagModifier) {
		return
	})
}

func (f *ConfigMustLoadFunc) nextHook() func(interface{}, ...tags.TagModifier) {
	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

// History returns a sequence of ConfigMustLoadFuncCall objects describing
// the invocations of this function.
func (f *ConfigMustLoadFunc) History() []ConfigMustLoadFuncCall {
	return f.history
}

// ConfigMustLoadFuncCall is an object that describes an invocation of
// method MustLoad on an instance of MockConfig.
type ConfigMustLoadFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 interface{}
	// Arg1 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg1 []tags.TagModifier
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c ConfigMustLoadFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg1 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigMustLoadFuncCall) Results() []interface{} {
	return []interface{}{}
}
