package utils

import (
	"github.com/iyear/biligo"
	"log"
	"math/rand"
	"time"
)

func GetRandomVideoAV(c *biligo.BiliClient) int64 {
	// TODO: have not implemented yet
	return 170001
	//return biligo.BV2AV("BV1764y197vL")
}

func WaitRandom() {
	wait := rand.Intn(10) + 5
	log.Printf("随机等待 %d s...\n\n", wait)
	time.Sleep(time.Duration(wait) * time.Second)
}
