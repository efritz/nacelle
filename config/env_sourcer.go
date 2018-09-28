package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	replacer = strings.NewReplacer(
		"\n", `\n`,
		"\t", `\t`,
		"\r", `\r`,
	)

	replacePattern = regexp.MustCompile(`[^A-Za-z0-9_]+`)
)

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

	return func(envTagValue string) (string, bool) {
		envTags := []string{
			strings.ToUpper(fmt.Sprintf("%s_%s", prefix, envTagValue)),
			strings.ToUpper(envTagValue),
		}

		for _, envTag := range envTags {
			if val, ok := os.LookupEnv(envTag); ok {
				return val, ok
			}
		}

		return "", false
	}
}
