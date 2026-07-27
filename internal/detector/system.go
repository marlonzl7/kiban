package detector

import (
	"fmt"
	"os/exec"
	"strings"
)

func DetectArchitecture() (string, error) {
	cmd := exec.Command("uname", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erro ao verificar arquitetura: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

func IsSudoAvailable() bool {
	_, err := exec.LookPath("sudo")
	return err == nil
}
