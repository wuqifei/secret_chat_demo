package server

import (
	"github.com/wuqifei/chat/protomodel"
	"github.com/wuqifei/server_lib/libio"
	"github.com/wuqifei/server_lib/libnet2"
)

// 交给handle处理
func Handle(sess libnet2.Session2Interface, val []byte) {

	protoB := val[:2]
	val = val[2:]

	p := libio.GetUint16BE(protoB)

	switch p {
	case protomodel.ProtoPing:
		{
			PingC2S(sess, val)
		}
	case protomodel.ProtoConfirm:
		{
			ConfirmC2S(sess, val)
		}
	case protomodel.ProtoSendSingleTextMsg:
		{
			SendSingleTextMsg(sess, val)
		}
	case protomodel.ProtoSendSinglePublishTextMsg:
		{
			SendSinglePublishTextMsg(sess, val)
		}
	}
}
