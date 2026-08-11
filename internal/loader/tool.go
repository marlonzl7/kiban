package loader

import (
	"fmt"
	"io/fs"

	"gopkg.in/yaml.v3"
)

type Tool struct {
	Name           string             `yaml:"name"`
	Description    string             `yaml:"description"`
	Category       string             `yaml:"category"`
	VersionFlag    string             `yaml:"version_flag"`
	Install        map[string]Install `yaml:"install"`
	Verify         Verify             `yaml:"verify"`
	PostInstall    PostInstall        `yaml:"post_install"`
	DefaultVersion string             `yaml:"default_version"`
}

type Install struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Command string `yaml:"cmd"`
	Sudo    bool   `yaml:"sudo"`
	Shell   bool   `yaml:"shell"`
}

type Verify struct {
	Command string `yaml:"cmd"`
	Expect  string `yaml:"expect"`
}

type PostInstall struct {
	Message string `yaml:"message"`
}

func LoadToolFile(fsys fs.FS, name string) (Tool, error) {
	data, err := fs.ReadFile(fsys, name)
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
