package main

import (
//	"fmt"
	"github.com/0xAX/notificator"
)
var notify *notificator.Notificator

func main() {
	notify = notificator.New(notificator.Options{
    DefaultIcon: "icon/default.png",
    AppName:     "My test App",
  })

  notify.Push("title", "Hey Aaarjan, Time for sleep", "/home/user/icon.png", notificator.UR_NORMAL)
}

