package config

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// NewTOMLFileSourcer creates a sourcer that reads configuration data
// form a single TOML file. For more details on loading from files, see
// the YAML file sourcer.
func NewTOMLFileSourcer(path string) (Sourcer, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read TOML config (%s)", err.Error())
	}

	values := map[string]interface{}{}
	if err := toml.Unmarshal(content, &values); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal TOML config (%s)", err.Error())
	}

	return newFileContentSourcer(values)
}
