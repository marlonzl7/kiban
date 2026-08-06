package cli

import (
	"fmt"

	"github.com/marlonzl7/kiban/internal/detector"
	"github.com/marlonzl7/kiban/internal/executor"
	"github.com/marlonzl7/kiban/internal/loader"
	"github.com/marlonzl7/kiban/internal/validator"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs tools defined in a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := detector.DetectDistro("/etc/os-release")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		packageManager, err := detector.DetectPackageManager(info)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		arch, err := detector.DetectArchitecture()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		sudoAvailable := detector.IsSudoAvailable()

		sudoAvailableMsg := ""

		if sudoAvailable {
			sudoAvailableMsg = "available"
		} else {
			sudoAvailableMsg = "not available"
		}

		fmt.Printf("Environment: %s (%s) | %s | sudo %s\n", info.ID, packageManager, arch, sudoAvailableMsg)

		setupFile, err := loader.LoadSetupFile("setup.yaml")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("setup file version %d loaded\n", setupFile.Version)

		tools, err := loader.LoadToolsFromDir("tools")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("supported tools loaded\n")

		errs := validator.Validate(setupFile, tools)
		if errs != nil {
			fmt.Printf("invalid setup file: the file structure contains errors\n")
			fmt.Println(errs.Error())
			return
		}

		fmt.Printf("valid setup. starting tool installation...\n")

		summary := executor.InstallAll(setupFile, tools, packageManager, arch)

		fmt.Printf("Summary: %d installed, %d failed\n", summary.Installed, summary.Failed)
	},
}
