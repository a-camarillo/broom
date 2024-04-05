package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
  Use: "run",
  Short: "Runs broom",
  Long: "Runs broom and opens up user interface for " +
        "cleaning up repository branches.",
  Run: func(cmd *cobra.Command, args []string) {
    InitializeMenu()
  },
}
