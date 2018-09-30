package config

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type YAMLFileSourcerSuite struct{}

func (s *YAMLFileSourcerSuite) TestLoadJSON(t sweet.T) {
	sourcer, err := NewYAMLFileSourcer("test-files/values.json")
	Expect(err).To(BeNil())

	ensureEquals(sourcer, "foo", "", "bar")
	ensureMatches(sourcer, "bar", "", "[1, 2, 3]")
	ensureMatches(sourcer, "baz", "", "null")
	ensureMatches(sourcer, "bonk", "", `{"x": 1, "y": 2, "z": 3}`)
	ensureMatches(sourcer, "encoded", "", `{"w": 4}`)
	ensureMatches(sourcer, "bonk", "x", "1")
	ensureMatches(sourcer, "encoded", "w", "4")
}

func (s *YAMLFileSourcerSuite) TestLoadYAML(t sweet.T) {
	sourcer, err := NewYAMLFileSourcer("test-files/values.yaml")
	Expect(err).To(BeNil())

	ensureEquals(sourcer, "foo", "", "bar")
	ensureMatches(sourcer, "bar", "", "[1, 2, 3]")
	ensureMatches(sourcer, "baz", "", "null")
	ensureMatches(sourcer, "bonk", "", `{"x": 1, "y": 2, "z": 3}`)
	ensureMatches(sourcer, "encoded", "", `{"w": 4}`)
	ensureMatches(sourcer, "bonk", "x", "1")
	ensureMatches(sourcer, "encoded", "w", "4")
}
