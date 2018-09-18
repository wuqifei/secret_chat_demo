package main

import (
	"github.com/wuqifei/chat/logs"
	"github.com/wuqifei/chat/server"
	"github.com/wuqifei/server_lib/signal"
)

func main() {

	logs.NewLog("./logs/chat.log")
	server.New(":10001")
	signal.InitSignal()
}
