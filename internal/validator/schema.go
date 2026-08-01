package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/marlonzl7/kiban/internal/loader"
)

func Validate(setupFile loader.SetupFile, tools []loader.Tool) error {
	nameToTool := make(map[string]loader.Tool)
	var errs []error

	for _, tool := range tools {
		nameToTool[tool.Name] = tool
	}

	for _, values := range setupFile.Tools {
		for _, value := range values {
			toolName, _, _ := strings.Cut(value, "@")

			if toolName == "" {
				errs = append(errs, fmt.Errorf("invalid tool entry %q: missing tool name", value))
				continue
			}

			_, exists := nameToTool[toolName]

			if !exists {
				errs = append(errs, fmt.Errorf("tool %s is not supported", toolName))
			}
		}
	}

	return errors.Join(errs...)
}
