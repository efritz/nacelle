package config

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type FileSourcerSuite struct{}

func (s *FileSourcerSuite) TestLoadJSON(t sweet.T) {
	sourcer, err := NewFileSourcer("test-files/values.json", ParseYAML)
	Expect(err).To(BeNil())

	ensureEquals(sourcer, []string{"foo"}, "bar")
	ensureMatches(sourcer, []string{"bar"}, "[1, 2, 3]")
	ensureMatches(sourcer, []string{"baz"}, "null")
	ensureMatches(sourcer, []string{"bonk"}, `{"x": 1, "y": 2, "z": 3}`)
	ensureMatches(sourcer, []string{"encoded"}, `{"w": 4}`)
	ensureMatches(sourcer, []string{"bonk.x"}, `1`)
	ensureMatches(sourcer, []string{"encoded.w"}, `4`)
	ensureMatches(sourcer, []string{"deeply.nested.struct"}, `[1, 2, 3]`)
}

func (s *FileSourcerSuite) TestLoadYAML(t sweet.T) {
	sourcer, err := NewFileSourcer("test-files/values.yaml", ParseYAML)
	Expect(err).To(BeNil())

	ensureEquals(sourcer, []string{"foo"}, "bar")
	ensureMatches(sourcer, []string{"bar"}, "[1, 2, 3]")
	ensureMatches(sourcer, []string{"baz"}, "null")
	ensureMatches(sourcer, []string{"bonk"}, `{"x": 1, "y": 2, "z": 3}`)
	ensureMatches(sourcer, []string{"encoded"}, `{"w": 4}`)
	ensureMatches(sourcer, []string{"bonk.x"}, `1`)
	ensureMatches(sourcer, []string{"encoded.w"}, `4`)
	ensureMatches(sourcer, []string{"deeply.nested.struct"}, `[1, 2, 3]`)
}

func (s *FileSourcerSuite) TestLoadTOML(t sweet.T) {
	sourcer, err := NewFileSourcer("test-files/values.toml", ParseTOML)
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
