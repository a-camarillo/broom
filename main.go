package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use: "broom",
	Short: "broom, sweep things",
	Long: `A longer description for what to do with a broom`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Broom")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}