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

		setupFile, err := loader.LoadSetupFile(setupFilePath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Setup file version %d loaded\n", setupFile.Version)

		tools, err := loader.LoadToolsFromDir(loader.ToolsFS, ".")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Supported tools loaded\n")

		errs := validator.ValidateSchema(setupFile, tools)
		if errs != nil {
			fmt.Printf("invalid setup file: the file structure contains errors\n")
			fmt.Println(errs.Error())
			return
		}

		errs = executor.ValidateVersionsMandatory(setupFile, tools, packageManager)
		if errs != nil {
			fmt.Printf("invalid setup file: the file structure contains errors\n")
			fmt.Println(errs.Error())
			return
		}

		fmt.Printf("Valid setup. starting tool installation...\n")

		err = executor.ValidateSudoSession()
		if err != nil {
			fmt.Println("unable to obtain sudo privileges")
			fmt.Println("check your password and verify whether your user has permission to use sudo")
			return
		}

		summary := executor.InstallAll(setupFile, tools, packageManager, arch)

		fmt.Printf("Summary: %d installed, %d failed\n", summary.Installed, summary.Failed)
	},
}
