package executor

import (
	"fmt"
	"strings"

	"github.com/marlonzl7/kiban/internal/loader"
)

func ResolveCommand(cmd string, params map[string]string) string {
	for key, value := range params {
		cmd = strings.ReplaceAll(cmd, "{{"+key+"}}", value)
	}

	return cmd
}

func ResolveVersion(tool loader.Tool, requestedVersion string, packageManager string) (string, error) {
	version := requestedVersion
	if version == "" {
		version = tool.DefaultVersion
	}

	install := tool.Install[packageManager]

	for _, step := range install.Steps {
		if strings.Contains(step.Command, "{{version}}") && version == "" {
			return "", fmt.Errorf("tool %s requires a version (use %s@<version>) but no version or default_version was found", tool.Name, tool.Name)
		}
	}

	return version, nil
}
