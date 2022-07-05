package run

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "运行",
	Long:  "执行bilibili每日任务",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var (
	DedeUserID      string
	DedeUserIDCkMd5 string
	SESSDATA        string
	BiliJCT         string
)

func init() {
	Cmd.Flags().StringVarP(&DedeUserID, "dedeuserid", "d", "", "登录用信息")
	Cmd.Flags().StringVarP(&DedeUserIDCkMd5, "md5", "m", "", "登录用信息")
	Cmd.Flags().StringVarP(&SESSDATA, "sessdata", "s", "", "登录用信息")
	Cmd.Flags().StringVarP(&BiliJCT, "bilijct", "j", "", "登录用信息")
}
