package config

import "github.com/aphistic/sweet"

type MultiSourcerSuite struct{}

func (s *MultiSourcerSuite) TestMultiSourcerBasic(t sweet.T) {
	s1, _ := NewMapSourcer(map[string]interface{}{"foo": "bar"})
	s2, _ := NewMapSourcer(map[string]interface{}{"bar": "baz"})
	s3, _ := NewMapSourcer(map[string]interface{}{"foo": "bonk"})
	multi := NewMultiSourcer(s1, s2, s3)

	ensureEquals(multi, "foo", "bar")
	ensureEquals(multi, "bar", "baz")
	ensureMissing(multi, "baz")
}
