package config

import "github.com/aphistic/sweet"

type MultiSourcerSuite struct{}

func (s *MultiSourcerSuite) TestMultiSourcerBasic(t sweet.T) {
	s1, _ := NewMapSourcer(map[string]interface{}{"foo": "bar"})
	s2, _ := NewMapSourcer(map[string]interface{}{"bar": "baz"})
	s3, _ := NewMapSourcer(map[string]interface{}{"foo": "bonk"})
	multi := NewMultiSourcer(s1, s2, s3)

	ensureEquals(multi, []string{"foo"}, "bar")
	ensureEquals(multi, []string{"bar"}, "baz")
	ensureMissing(multi, []string{"baz"})
}

func (s *MultiSourcerSuite) TestMultiSourcerContext(t sweet.T) {
	s1, _ := NewMapSourcer(map[string]interface{}{"x": `{"y": {"w": 42}}`})
	s2, _ := NewMapSourcer(map[string]interface{}{"x": `{"y": {"z": 42}}`})
	s3, _ := NewMapSourcer(map[string]interface{}{"x": `{"y": {"z": "foo"}}`})
	multi := NewMultiSourcer(s1, s2, s3)

	ensureEquals(multi, []string{"x", "y", "z"}, "42")
}
