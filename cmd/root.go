package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "broom",
	Version: "0.3.0",
	Short:   "Broom sweeps up your local repository branches",
	Long: "A terminal user interface tool for cleaning up your repository's local branches.\n" +
		"Includes options for viewing remote references.\n" +
		"Documentation can be found at https://pkg.go.dev/github.com/a-camarillo/broom.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
