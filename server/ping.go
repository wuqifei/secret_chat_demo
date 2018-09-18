package server

import (
	"time"

	"github.com/wuqifei/chat/packet"

	"github.com/wuqifei/chat/protomodel"
	"github.com/wuqifei/server_lib/libnet2"
)

func PingC2S(sess libnet2.Session2Interface, val []byte) {
	ret := &protomodel.PingS2C{}
	ret.TimeStamp = time.Now().Unix()
	retB := packet.RollPacket(protomodel.ProtoPing, ret)
	sess.Send(retB)
}
