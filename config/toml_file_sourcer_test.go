package config

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type TOMLFileSourcerSuite struct{}

func (s *TOMLFileSourcerSuite) TestLoadTOML(t sweet.T) {
	sourcer, err := NewTOMLFileSourcer("test-files/values.toml")
	Expect(err).To(BeNil())

	ensureEquals(sourcer, "foo", "", "bar")
	ensureMatches(sourcer, "bar", "", "[1, 2, 3]")
	ensureMissing(sourcer, "baz", "")
	ensureMatches(sourcer, "bonk", "", `{"x": 1, "y": 2, "z": 3}`)
	ensureMatches(sourcer, "encoded", "", `{"w": 4}`)
	ensureMatches(sourcer, "bonk", "x", "1")
	ensureMatches(sourcer, "encoded", "w", "4")
}
