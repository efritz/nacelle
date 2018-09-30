package config

import (
	"testing"

	"github.com/aphistic/sweet"
	"github.com/aphistic/sweet-junit"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	RegisterFailHandler(sweet.GomegaFail)

	sweet.Run(m, func(s *sweet.S) {
		s.RegisterPlugin(junit.NewPlugin())

		s.AddSuite(&ConfigSuite{})
		s.AddSuite(&EnvSourcerSuite{})
		s.AddSuite(&JSONSuite{})
		s.AddSuite(&LoggingConfigSuite{})
		s.AddSuite(&MapSourcerSuite{})
		s.AddSuite(&MultiSourcerSuite{})
		s.AddSuite(&TOMLFileSourcerSuite{})
		s.AddSuite(&YAMLFileSourcerSuite{})
	})
}

//
//

func ensureEquals(sourcer Sourcer, key, context, expected string) {
	val, ok := sourcer(key, context)
	Expect(ok).To(BeTrue())
	Expect(val).To(Equal(expected))
}

func ensureMatches(sourcer Sourcer, key, context, expected string) {
	val, ok := sourcer(key, context)
	Expect(ok).To(BeTrue())
	Expect(val).To(MatchJSON(expected))
}

func ensureMissing(sourcer Sourcer, key, context string) {
	_, ok := sourcer(key, context)
	Expect(ok).To(BeFalse())
}
