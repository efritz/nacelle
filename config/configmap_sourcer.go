package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// NewConfigMapSourcer creates a sourcer that reads from a
// Kubernetes ConfigMap volume mounted into a given directory.
// The content of the files are assumed to be raw strings or
// JSON encoded values which will be decoded at the time a
// config struct is populated.
func NewConfigMapSourcer(directory string) (Sourcer, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("failed to read configmap directory (%s)", err.Error())
	}

	contentMap := map[string]string{}
	for _, file := range files {
		content, err := ioutil.ReadFile(filepath.Join(directory, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read configmap file '%s' (%s)", file.Name(), err.Error())
		}

		contentMap[file.Name()] = string(content)
	}

	return func(envTag string) (string, bool) {
		val, ok := contentMap[envTag]
		return val, ok
	}, nil
}
