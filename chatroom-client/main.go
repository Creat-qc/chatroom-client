package main

import (
	"chatroom-client/chat"
	"fmt"
	"github.com/liangdas/armyant/task"
	"os"
	"os/signal"
)

func main() {
	task := task.LoopTask{
		C: 1,
	}
	manager := chat.NewManger(task)
	fmt.Println("chatroom-client启动")
	task.Run(manager)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	task.Stop()
	os.Exit(1)
}
