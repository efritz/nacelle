package config

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type JSONSuite struct{}

func (s *JSONSuite) TestToJSONString(t sweet.T) {
	var val string
	ok := toJSON([]byte("foobar"), &val)
	Expect(ok).To(BeTrue())
	Expect(val).To(Equal("foobar"))
}

func (s *JSONSuite) TestToJSONNonString(t sweet.T) {
	var val []int
	ok := toJSON([]byte("[1, 2, 3, 4, 5]"), &val)
	Expect(ok).To(BeTrue())
	Expect(val).To(Equal([]int{1, 2, 3, 4, 5}))
}

func (s *JSONSuite) TestToJSONBadType(t sweet.T) {
	var val []int
	ok := toJSON([]byte(`[1, 2, "3", 4, 5]`), &val)
	Expect(ok).To(BeFalse())
}

func (s *JSONSuite) TestQuoteJSON(t sweet.T) {
	json := quoteJSON([]byte(`
	foo
	bar
	baz`))

	Expect(json).To(MatchJSON(`"\n\tfoo\n\tbar\n\tbaz"`))
}

func (s *JSONSuite) TestExtractContext(t sweet.T) {
	inner, ok := extractContext(`{"points": {"x": 1, "y": 2, "z": 3}}`, "points")
	Expect(ok).To(BeTrue())
	Expect(inner).To(MatchJSON(`{"x": 1, "y": 2, "z": 3}`))
}

func (s *JSONSuite) TestExtractContextNested(t sweet.T) {
	inner, ok := extractContext(`{"a": {"b": {"c": {"x": 1, "y": 2, "z": 3}}}}`, "a.b.c")
	Expect(ok).To(BeTrue())
	Expect(inner).To(MatchJSON(`{"x": 1, "y": 2, "z": 3}`))
}

func (s *JSONSuite) TestExtractContextMissing(t sweet.T) {
	_, ok := extractContext(`{"b": {"x": 1, "y": 2, "z": 3}}`, "a")
	Expect(ok).To(BeFalse())
}
func (s *JSONSuite) TestExtractContextNoContext(t sweet.T) {
	inner, ok := extractContext(` {"x": 1, "y": 2, "z": 3}`, "")
	Expect(ok).To(BeTrue())
	Expect(inner).To(MatchJSON(`{"x": 1, "y": 2, "z": 3}`))
}
