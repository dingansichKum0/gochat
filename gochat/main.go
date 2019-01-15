package main

import (
	"fmt"
	"time"

	"github.com/Zhan9Yunhua/gochat"
)

func main() {
	c, _ := gochat.NewChat(&gochat.Chat{
		QrcodeProt:   8005,
		IsQrcodeFile: false,
	})
	c.Start()
	for {
		c.SendMessage("YinT1e Ji", fmt.Sprintf("* test  send a message to YinT1e Ji* %s", time.Now()))
		time.Sleep(time.Second * 60)
	}
}
