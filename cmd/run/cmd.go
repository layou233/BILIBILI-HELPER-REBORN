package run

import (
	"BILIBILI-HELPER-REBORN/task/exp"
	"BILIBILI-HELPER-REBORN/task/videos"
	"BILIBILI-HELPER-REBORN/utils"
	"github.com/iyear/biligo"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "运行",
	Long:  "执行bilibili每日任务",
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

	retriedTimes := 0
watch:
	utils.WaitRandom()
	log.Println("=====每日视频观看开始=====")
	videoInfo, playedTime, err := videos.DoVideoWatch(client)
	if err != nil {
		log.Println("失败，原因:", err)
		log.Println("=====每日视频观看失败=====\n")
		if retriedTimes < 5 {
			retriedTimes++
			goto watch
		} else {
			log.Println("每日视频观看 失败超过5次，不再重试，请报告作者!")
		}
	} else {
		log.Println("已播放视频:", videoInfo.Title, ", "+videoInfo.BVID, ", av"+strconv.FormatInt(videoInfo.AID, 10))
		log.Println("播放到:", playedTime, "s")
		log.Println("=====每日视频观看结束=====\n")
	}

	retriedTimes = 0
share:
	utils.WaitRandom()
	log.Println("=====每日视频分享开始=====")
	err = videos.DoVideoShare(client)
	if err != nil {
		log.Println("失败，原因:", err)
		log.Println("=====每日视频分享失败=====")
		if retriedTimes < 5 {
			retriedTimes += 1
			goto share
		} else {
			log.Println("每日视频分享 失败超过5次，不再重试，请报告作者!")
		}
	} else {
		log.Println("分享成功!")
		log.Println("=====每日视频分享结束=====\n")
	}

	utils.WaitRandom()
	reward, err = exp.GetDailyStatus(client)
	if err != nil {
		log.Panic("error when requesting user info:", err)
		return
	}
	exp.PrintDailyStatus(reward)
}
