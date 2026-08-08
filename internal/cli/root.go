package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kiban",
	Short: "CLI for setting up development environments",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kiban - Development environment setup CLI")
	},
}

var setupFilePath string

func init() {
	installCmd.Flags().StringVar(&setupFilePath, "file", "setup.yaml", "path to the setup.yaml file")
	rootCmd.AddCommand(installCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
