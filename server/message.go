package server

import (
	"github.com/gogo/protobuf/proto"
	"github.com/wuqifei/chat/logs"
	"github.com/wuqifei/chat/packet"
	"github.com/wuqifei/chat/protomodel"
	"github.com/wuqifei/server_lib/libencrypt/aes_ecb_encrypt"
	"github.com/wuqifei/server_lib/libencrypt/base64_encrypt"
	"github.com/wuqifei/server_lib/libnet2"
)

func SendSingleTextMsg(sess libnet2.Session2Interface, val []byte) {

	m := new(protomodel.SendSingleTextMsgC2S)
	ret := new(protomodel.SendSingleTextMsgS2C)
	err := proto.Unmarshal(val, m)

	if err != nil {
		logs.Error("unmarshal error [%v]", err)
		ret.Flag = false
		retB := packet.RollPacket(protomodel.ProtoSendSingleTextMsg, ret)
		sess.Send(retB)
		return
	}
	// 需要base64，解码一次,防止数据过大
	decodeMsg, err := base64_encrypt.Base64StdDecode(m.Message)
	if err != nil {
		logs.Error("decodeMsg error [%v]", err)
		ret.Flag = false
		retB := packet.RollPacket(protomodel.ProtoSendSingleTextMsg, ret)
		sess.Send(retB)
		return
	}
	// 这里需要用aes给消息加密

	key, _ := sess.Get("random_key")
	enVal := aes_ecb_encrypt.Encrypt(decodeMsg, key.(string))
	encodeMsg := base64_encrypt.Base64StdEncode(enVal)
	if err != nil {
		logs.Error("base64_encrypt error [%v]", err)
		ret.Flag = false
		retB := packet.RollPacket(protomodel.ProtoSendSingleTextMsg, ret)
		sess.Send(retB)
		return
	}

	pmsg := new(protomodel.SendSinglePublishTextMsgS2C)
	pmsg.Uid = sess.GetUniqueID()
	pmsg.Message = encodeMsg

	sessRet := talkPool.Get(m.Uid)
	if sessRet == nil {
		ret.Flag = false
		retB := packet.RollPacket(protomodel.ProtoSendSinglePublishTextMsg, ret)
		sess.Send(retB)
		return
	}

	// 发送消息
	sess2 := sessRet.(libnet2.Session2Interface)
	retBP := packet.RollPacket(protomodel.ProtoSendSinglePublishTextMsg, pmsg)
	sess2.Send(retBP)

	// 发送消息
	ret.Flag = true
	retB := packet.RollPacket(protomodel.ProtoSendSingleTextMsg, ret)
	sess.Send(retB)
}

func SendSinglePublishTextMsg(sess libnet2.Session2Interface, val []byte) {
	// 接收到数据
}
