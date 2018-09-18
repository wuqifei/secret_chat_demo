package main

import (
	"fmt"
	"net"

	"github.com/wuqifei/chat/client/cpacket"
)

func main() {
	conn, err := net.Dial("tcp", "198.13.48.130:10001")
	if err != nil {
		panic(err)
	}
	fmt.Println("已连接服务器")
	defer conn.Close()

	go cpacket.ReadMsg(conn)

	cpacket.UIInit()
}
