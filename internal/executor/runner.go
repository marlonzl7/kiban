package executor

import (
	"fmt"
	"os/exec"

	"github.com/marlonzl7/kiban/internal/loader"
	"github.com/marlonzl7/kiban/internal/utils"
)

func RunStep(step loader.Step) error {
	args := utils.BuildArgs(step.Command, step.Sudo)

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("error executing step %q: %w\noutput: %s", step.Command, err, output)
	}

	return nil
}

func RunSteps(steps []loader.Step, version string, arch string) error {
	params := map[string]string{
		"version": version,
		"arch":    arch,
	}

	for _, step := range steps {
		step.Command = ResolveCommand(step.Command, params)

		err := RunStep(step)
		if err != nil {
			return err
		}
	}

	return nil
}
