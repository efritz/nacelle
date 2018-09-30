package config

import (
	"encoding/json"
	"fmt"
)

// NewMapSourcer creates a sourcer that reads keys from the given map of
// interface values. Partially due to the way defaults are implemented,
// all config map values must be JSON-serializable - for configuration
// data, this is a reasonable requirement. This requirement also extends
// to the sourcers in this package which read from files.
func NewMapSourcer(values map[string]interface{}) (Sourcer, error) {
	jsonValues := map[string]string{}
	for key, value := range values {
		serialized, err := serializeJSONValue(value)
		if err != nil {
			return nil, fmt.Errorf("illegal configuration value for '%s' (%s)", key, err.Error())
		}

		jsonValues[key] = serialized
	}

	return func(env, context string) (string, bool) {
		val, ok := jsonValues[env]
		if !ok {
			return "", false
		}

		return extractContext(val, context)
	}, nil
}

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
