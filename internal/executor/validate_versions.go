package executor

import (
	"errors"
	"strings"

	"github.com/marlonzl7/kiban/internal/loader"
	"github.com/marlonzl7/kiban/internal/utils"
)

func ValidateVersionsMandatory(setupFile loader.SetupFile, tools []loader.Tool, packageManager string) error {
	var errs []error
	nameToTool := utils.NameToTool(tools)

	for _, values := range setupFile.Tools {
		for _, value := range values {
			toolName, toolVersion, _ := strings.Cut(value, "@")

			tool := nameToTool[toolName]

			_, err := ResolveVersion(tool, toolVersion, packageManager)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errors.Join(errs...)
}
