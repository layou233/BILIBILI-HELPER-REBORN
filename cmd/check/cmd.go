package check

import (
	"BILIBILI-HELPER-REBORN/task/exp"
	"BILIBILI-HELPER-REBORN/utils"
	"github.com/spf13/cobra"
	"log"
)
import "github.com/iyear/biligo"

var Cmd = &cobra.Command{
	Use:   "check",
	Short: "检查",
	Long:  "检查账号可用性并打印信息",
	Run:   run,
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

func run(cmd *cobra.Command, args []string) {
	client, err := biligo.NewBiliClient(&biligo.BiliSetting{
		Auth: &biligo.CookieAuth{
			DedeUserID:      DedeUserID,
			DedeUserIDCkMd5: DedeUserIDCkMd5,
			SESSDATA:        SESSDATA,
			BiliJCT:         BiliJCT,
		},
	})
	if err != nil {
		log.Panic("unknown fatal error:", err)
		return
	}

	userInfo, err := client.UserGetInfo(client.Me.MID)
	if err != nil {
		log.Panic("error when requesting user info:", err)
		return
	}
	reward, err := exp.GetDailyStatus(client)
	if err != nil {
		log.Panic("error when requesting user info:", err)
		return
	}
	vipStat, err := client.GetVipStat()
	if err != nil {
		log.Panic("error when requesting user info:", err)
		return
	}

	log.Printf(`
=====登录成功=====
| 昵称: %v
| 硬币数: %v, 今日投币已获得经验: %v
| 大会员: %v
================`, client.Me.UName, userInfo.Coins, reward.Coins,
		utils.GetVIPStat(vipStat).String())

	exp.PrintDailyStatus(reward)
}
