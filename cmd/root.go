package cmd

import (
	"BILIBILI-HELPER-REBORN/cmd/check"
	"BILIBILI-HELPER-REBORN/cmd/run"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(run.Cmd)
	rootCmd.AddCommand(check.Cmd)
}
