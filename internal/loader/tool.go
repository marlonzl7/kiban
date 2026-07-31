package loader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Tool struct {
	Name        string             `yaml:"name"`
	Description string             `yaml:"description"`
	Category    string             `yaml:"category"`
	VersionFlag string             `yaml:"version_flag"`
	Install     map[string]Install `yaml:"install"`
	Verify      Verify             `yaml:"verify"`
	PostInstall PostInstall        `yaml:"post_install"`
}

type Install struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Command string `yaml:"cmd"`
	Sudo    bool   `yaml:"sudo"`
}

type Verify struct {
	Command string `yaml:"cmd"`
	Expect  string `yaml:"expect"`
}

type PostInstall struct {
	Message string `yaml:"message"`
}

func LoadToolFile(path string) (Tool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Tool{}, fmt.Errorf("failed to load tool files: %w", err)
	}

	var tool Tool

	err = yaml.Unmarshal(data, &tool)
	if err != nil {
		return Tool{}, fmt.Errorf("failed to deserialize tool files: %w", err)
	}

	return tool, nil
}
