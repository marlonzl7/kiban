package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kiban",
	Short: "CLI for setting up development environments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kiban - Development environment setup CLI")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
