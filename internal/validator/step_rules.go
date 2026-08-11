package validator

import (
	"errors"
	"fmt"

	"github.com/marlonzl7/kiban/internal/loader"
)

func ValidateStepShellSudoExclusivity(tools []loader.Tool) error {
	var errs []error

	for _, tool := range tools {
		for pm, install := range tool.Install {
			for _, step := range install.Steps {
				if step.Shell && step.Sudo {
					errs = append(errs, fmt.Errorf(
						"tool %s (%s): step %q cannot set both shell and sudo — "+
							"use sudo explicitly inside cmd when shell is true",
						tool.Name, pm, step.Command,
					))
				}
			}
		}
	}

	return errors.Join(errs...)
}