package videos

import (
	"BILIBILI-HELPER-REBORN/utils"
	"github.com/iyear/biligo"
	"math/rand"
)

func DoVideoShare(c *biligo.BiliClient) (err error) {
	_, err = c.VideoShare(utils.GetRandomVideoAV(c))
	return
}

func DoVideoWatch(c *biligo.BiliClient) (info *biligo.VideoInfo, playedTime int64, err error) {
	//new:
	av := utils.GetRandomVideoAV(c)
	info, _ = c.VideoGetInfo(av)
	/*if info.Videos != 1 {
		// 因为如果是多P的稿件，获取第一P的视频时长比较困难
		// 换句话说就是我比较懒，这里对多P视频的UP说声抱歉
		goto new
	}*/
	playedTime = rand.Int63n(50) + 1
	//playedTime:=rand.Int63n(info.Duration-1)+1 // 这里因为此随机函数取值范围为[0,n)，故手动+1
	err = c.VideoHeartBeat(av, 0, // 只播放第一P
		playedTime)
	return
}
