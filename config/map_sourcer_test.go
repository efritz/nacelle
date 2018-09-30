package config

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type MapSourcerSuite struct{}

func (s *MapSourcerSuite) TestValues(t sweet.T) {
	values := map[string]interface{}{
		"foo": "bar",
		"bar": []int{1, 2, 3},
		"baz": nil,
		"bonk": map[string]int{
			"x": 1,
			"y": 2,
			"z": 3,
		},
		"encoded": `{"w": 4}`,
	}

	sourcer, err := NewMapSourcer(values)
	Expect(err).To(BeNil())

	ensureEquals(sourcer, "foo", "", "bar")
	ensureMatches(sourcer, "bar", "", "[1, 2, 3]")
	ensureMatches(sourcer, "baz", "", "null")
	ensureMatches(sourcer, "bonk", "", `{"x": 1, "y": 2, "z": 3}`)
	ensureMatches(sourcer, "encoded", "", `{"w": 4}`)
	ensureMatches(sourcer, "bonk", "x", "1")
	ensureMatches(sourcer, "encoded", "w", "4")
}

func (s *MapSourcerSuite) TestIllegalValue(t sweet.T) {
	_, err := NewMapSourcer(map[string]interface{}{
		"double": func(x int) int { return x * 2 },
	})

	Expect(err).To(MatchError("illegal configuration value for 'double' (json: unsupported type: func(int) int)"))
}
