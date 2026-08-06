package executor

import (
	"fmt"
	"strings"

	"github.com/marlonzl7/kiban/internal/loader"
	"github.com/marlonzl7/kiban/internal/utils"
)

type Summary struct {
	Installed int
	Failed    int
}

func InstallAll(setupFile loader.SetupFile, tools []loader.Tool, packageManager, arch string) Summary {
	var summary Summary
	nameToTool := utils.NameToTool(tools)

	for _, values := range setupFile.Tools {
		for _, value := range values {
			toolName, toolVersion, _ := strings.Cut(value, "@")

			fmt.Printf("Installing %s... ", toolName)

			tool, exists := nameToTool[toolName]
			if !exists {
				fmt.Printf("[FAIL] ")
				fmt.Printf("(tool %s is not supported)\n", toolName)
				summary.Failed++
				continue
			}

			steps := tool.Install[packageManager].Steps
			if steps == nil {
				fmt.Printf("[FAIL] ")
				fmt.Printf("(package manager %s is not supported)\n", toolName)
				summary.Failed++
				continue
			}

			err := RunSteps(steps, toolVersion, arch)
			if err != nil {
				fmt.Printf("[FAIL] ")
				fmt.Printf("(%s)\n", err)
				summary.Failed++
				continue
			}

			err = VerifyInstallation(tool.Verify)
			if err != nil {
				fmt.Printf("[FAIL] ")
				fmt.Printf("(%s)\n", err)
				summary.Failed++
				continue
			}

			fmt.Printf("[OK]\n")

			summary.Installed++
		}
	}

	return summary
}
