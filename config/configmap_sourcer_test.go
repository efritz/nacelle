package config

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type ConfigMapSourcerSuite struct{}

func (s *ConfigMapSourcerSuite) TestValues(t sweet.T) {
	sourcer, err := NewConfigMapSourcer("test-files/k8s")
	Expect(err).To(BeNil())

	ensureEquals(sourcer, "a", "content of a")
	ensureEquals(sourcer, "b", "content of b")
	ensureEquals(sourcer, "c", "content of c")
	ensureEquals(sourcer, "x.y.z", "content of xyz")
	ensureMissing(sourcer, "d")
}
