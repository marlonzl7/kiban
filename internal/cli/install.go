package cli

import (
	"fmt"

	"github.com/marlonzl7/kiban/internal/detector"
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

		fmt.Println("OS:", info.ID)
		fmt.Println("VERSION:", info.Version)
		fmt.Println("LIKE:", info.IDLike)
		fmt.Println("PACKAGE_MANAGER:", packageManager)
		fmt.Println("ARCHITECTURE:", arch)
		fmt.Println("IS_SUDO_AVAILABLE:", sudoAvailable)
	},
}
