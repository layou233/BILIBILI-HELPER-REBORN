package main

import (
	"BILIBILI-HELPER-REBORN/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Panicln("cmd.Execute err:", err)
	}
}
