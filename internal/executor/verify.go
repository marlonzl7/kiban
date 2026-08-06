package executor

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/marlonzl7/kiban/internal/loader"
	"github.com/marlonzl7/kiban/internal/utils"
)

func VerifyInstallation(verify loader.Verify) error {
	args := utils.BuildArgs(verify.Command, false)

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("error executing verification command %q: %w\noutput: %s", verify.Command, err, output)
	}

	if !strings.Contains(string(output), verify.Expect) {
		return fmt.Errorf("verification failed for command %q: expected output to contain %q, got: %s", verify.Command, verify.Expect, output)
	}

	return nil
}
