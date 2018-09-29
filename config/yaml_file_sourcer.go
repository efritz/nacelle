package config

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// NewYAMLFileSourcer creates a sourcer that reads configuration data
// form a single file. It is expected that the top-level structure in
// this file is a mapping from configuration key names to their values.
// Values may be JSON-encoded strings, in which case they will be decoded
// implicitly. As YAML is a superset of JSON, this source can also parse
// JSON files. If you need to read from multiple files, use a MultiSourcer
// with multiple instances of this sourcer.
func NewYAMLFileSourcer(path string) (Sourcer, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML config (%s)", err.Error())
	}

	values := map[string]interface{}{}
	if err := yaml.Unmarshal(content, &values); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal YAML config (%s)", err.Error())
	}

	return NewMapSourcer(values)
}
