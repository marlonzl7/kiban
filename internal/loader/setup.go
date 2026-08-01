package loader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type SetupFile struct {
	Version int                 `yaml:"version"`
	Tools   map[string][]string `yaml:"tools"`
}

func LoadSetupFile(path string) (SetupFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return SetupFile{}, fmt.Errorf("failed to load setup file: %w", err)
	}

	var sf SetupFile

	err = yaml.Unmarshal(data, &sf)
	if err != nil {
		return SetupFile{}, fmt.Errorf("failed to deserialize setup file: %w", err)
	}

	return sf, nil
}
