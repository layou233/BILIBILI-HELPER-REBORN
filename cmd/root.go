package cmd

import (
	"BILIBILI-HELPER-REBORN/cmd/check"
	"BILIBILI-HELPER-REBORN/cmd/run"
	"BILIBILI-HELPER-REBORN/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Long:    "BILIBILI-HELPER 重制版，基于 iyear/biligo，使用Go编写.",
	Short:   "BiliBili每日任务助手",
	Version: version.Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(run.Cmd)
	rootCmd.AddCommand(check.Cmd)
}
