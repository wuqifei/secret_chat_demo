package server

import (
	"time"

	"github.com/wuqifei/chat/packet"
	"github.com/wuqifei/chat/protomodel"
	"github.com/wuqifei/server_lib/libnet2"
	"github.com/wuqifei/server_lib/librand"
)

func ConfirmS2C(sessId uint64, sess libnet2.Session2Interface) {
	AddConnSess(sess)
	randomKey := librand.CreateASCIIRandomCode(64)
	sess.Set("random_key", randomKey)
	sess.Set("conn_time", time.Now().UTC().Unix())
	m := new(protomodel.ConfirmS2C)
	m.RandomKey = randomKey
	m.ClientId = int64(sess.GetUniqueID())
	val := packet.RollPacket(protomodel.ProtoConfirm, m)
	sess.Send(val)
}

func ConfirmC2S(sess libnet2.Session2Interface, val []byte) {
	// 接收验证
	AddTalkSess(sess)
}
