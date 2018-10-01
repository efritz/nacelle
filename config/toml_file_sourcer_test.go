package config

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type TOMLFileSourcerSuite struct{}

func (s *TOMLFileSourcerSuite) TestLoadTOML(t sweet.T) {
	sourcer, err := NewTOMLFileSourcer("test-files/values.toml")
	Expect(err).To(BeNil())

	ensureEquals(sourcer, []string{"foo"}, "bar")
	ensureMatches(sourcer, []string{"bar"}, "[1, 2, 3]")
	ensureMissing(sourcer, []string{"baz"})
	ensureMatches(sourcer, []string{"bonk"}, `{"x": 1, "y": 2, "z": 3}`)
	ensureMatches(sourcer, []string{"encoded"}, `{"w": 4}`)
	ensureMatches(sourcer, []string{"bonk.x"}, `1`)
	ensureMatches(sourcer, []string{"encoded.w"}, `4`)
	ensureMatches(sourcer, []string{"deeply.nested.struct"}, `[1, 2, 3]`)
}
