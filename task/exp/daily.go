package exp

import (
	"github.com/iyear/biligo"
	"log"
)

func GetDailyStatus(c *biligo.BiliClient) (*biligo.ExpRewardStat, error) {
	reward, err := c.GetExpRewardStat()
	if err != nil {
		return nil, err
	}
	reward.Coins, err = c.GetExpCoinReward()
	if err != nil {
		return nil, err
	}
	return reward, nil
}

func PrintDailyStatus(reward *biligo.ExpRewardStat) {
	log.Println(`
=====每日任务完成情况=====
| { true 为 已完成, false 为 未完成 }
| 每日登录:`, reward.Login, `
| 每日观看视频:`, reward.Watch, `
| 每日分享:`, reward.Share, `
=======================`+"\n")
}
