package cmd

import (
	"github.com/a-camarillo/broom/ui"
	"github.com/spf13/cobra"
)

var Remotes bool

func init() {
  rootCmd.AddCommand(runCmd)
  runCmd.Flags().BoolVarP(&Remotes,"with-remotes","w",false,"Additionally show remote branches")
}

var runCmd = &cobra.Command{
  Use: "run",
  Short: "Runs broom",
  Long: "Runs broom and opens up user interface for " +
        "cleaning up repository branches.",
  Run: func(cmd *cobra.Command, args []string) {
    ui.Initialize(Remotes)
  },
}

