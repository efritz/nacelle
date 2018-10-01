package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

type fileContentSourcer struct {
	values map[string]string
}

func newFileContentSourcer(values map[string]interface{}) (Sourcer, error) {
	jsonValues := map[string]string{}
	for key, value := range values {
		serialized, err := serializeJSONValue(value)
		if err != nil {
			return nil, fmt.Errorf("illegal configuration value for '%s' (%s)", key, err.Error())
		}

		jsonValues[key] = serialized
	}

	return &fileContentSourcer{values: jsonValues}, nil
}

func (s *fileContentSourcer) Tags() []string {
	return []string{"file"}
}

func (s *fileContentSourcer) Get(values []string) (string, bool, bool) {
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
