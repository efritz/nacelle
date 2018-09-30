package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

var replacer = strings.NewReplacer(
	"\n", `\n`,
	"\t", `\t`,
	"\r", `\r`,
)

func toJSON(data []byte, v interface{}) bool {
	if json.Unmarshal(data, v) == nil {
		return true
	}

	ptr := reflect.ValueOf(v)

	if ptr.Kind() == reflect.Ptr && reflect.Indirect(ptr).Kind() == reflect.String {
		if json.Unmarshal(quoteJSON(data), v) == nil {
			return true
		}
	}

	return false
}

func quoteJSON(data []byte) []byte {
	return []byte(fmt.Sprintf(`"%s"`, replacer.Replace(string(data))))
}

func extractContext(val, context string) (string, bool) {
	if context == "" {
		return val, true
	}

	for _, part := range strings.Split(context, ".") {
		mapping := map[string]json.RawMessage{}
		if err := json.Unmarshal([]byte(val), &mapping); err != nil {
			return "", false
		}

		inner, ok := mapping[part]
		if !ok {
			return "", false
		}

		val = string(inner)
	}

	return val, true
}
