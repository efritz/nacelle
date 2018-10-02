package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ghodss/yaml"
)

type (
	fileSourcer struct {
		values map[string]string
	}

	FileParser func(content []byte) (map[string]interface{}, error)
)

var ParserMap = map[string]FileParser{
	"yaml": ParseYAML,
	"yml":  ParseYAML,
	"json": ParseYAML,
	"toml": ParseTOML,
}

// ParseYAML parses the given content as YAML.
func ParseYAML(content []byte) (map[string]interface{}, error) {
	values := map[string]interface{}{}
	if err := yaml.Unmarshal(content, &values); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal YAML config (%s)", err.Error())
	}

	return values, nil
}

// ParseTOML parses the given content as JSON.
func ParseTOML(content []byte) (map[string]interface{}, error) {
	values := map[string]interface{}{}
	if err := toml.Unmarshal(content, &values); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal TOML config (%s)", err.Error())
	}

	return values, nil
}

// NewYAMLFileSourcer creates a file sourcer that parses conent as YAML.
func NewYAMLFileSourcer(filename string) (Sourcer, error) {
	return NewFileSourcer(filename, ParseYAML)
}

// NewJSONFileSourcer creates a file sourcer that parses conent as TOML.
func NewTOMLFileSourcer(filename string) (Sourcer, error) {
	return NewFileSourcer(filename, ParseTOML)
}

// NewFileSourcer creates a sourcer that reads content from a file. The format
// of the file is read by the given FileParser. The content of the file must be
// an encoding of a map from string keys to JSON-serializable values.
func NewFileSourcer(filename string, parser FileParser) (Sourcer, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file '%s' (%s)", filename, err.Error())
	}

	values, err := parser(content)
	if err != nil {
		return nil, err
	}

	jsonValues := map[string]string{}
	for key, value := range values {
		serialized, err := serializeJSONValue(value)
		if err != nil {
			return nil, fmt.Errorf("illegal configuration value for '%s' (%s)", key, err.Error())
		}

		jsonValues[key] = serialized
	}

	return &fileSourcer{values: jsonValues}, nil
}

func (s *fileSourcer) Tags() []string {
	return []string{"file"}
}

func (s *fileSourcer) Get(values []string) (string, bool, bool) {
	if values[0] == "" {
		return "", true, false
	}

	segments := strings.Split(values[0], ".")

	if val, ok := s.values[segments[0]]; ok {
		if val, ok := extractJSONPath(val, segments[1:]); ok {
			return val, false, true
		}
	}

	return "", false, false
}

//
// Helpers

func serializeJSONValue(value interface{}) (string, error) {
	if str, ok := value.(string); ok {
		return str, nil
	}

	serialized, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(serialized), nil
}

func extractJSONPath(val string, path []string) (string, bool) {
	if len(path) == 0 {
		return val, true
	}

	for _, segment := range path {
		mapping := map[string]json.RawMessage{}
		if err := json.Unmarshal([]byte(val), &mapping); err != nil {
			return "", false
		}

		inner, ok := mapping[segment]
		if !ok {
			return "", false
		}

		val = string(inner)
	}

	return val, true
}
