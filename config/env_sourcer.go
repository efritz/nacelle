package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var replacePattern = regexp.MustCompile(`[^A-Za-z0-9_]+`)

// NewEnvSourcer creates a Sourcer that pulls values from the environment.
// the {PREFIX}_{NAME} envvar is read before falling back to the {NAME} envvar.
// The prefix will be normalized (replaces all non-alpha characters with an
// underscore and trims leading, trailing, and collapses consecutive underscores).
func NewEnvSourcer(prefix string) Sourcer {
	prefix = strings.Trim(
		string(replacePattern.ReplaceAll(
			[]byte(prefix),
			[]byte("_"),
		)),
		"_",
	)

	return func(path string) (string, bool) {
		envvars := []string{
			strings.ToUpper(fmt.Sprintf("%s_%s", prefix, path)),
			strings.ToUpper(path),
		}

		for _, envvar := range envvars {
			if val, ok := os.LookupEnv(envvar); ok {
				return val, true
			}
		}

		return "", false
	}
}
