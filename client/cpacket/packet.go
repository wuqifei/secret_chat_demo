package cpacket

import (
	"encoding/json"
	"io"
	"net"

	"github.com/wuqifei/server_lib/libencrypt/base64_encrypt"
	"github.com/wuqifei/server_lib/libfile"

	"github.com/wuqifei/server_lib/libencrypt/rsa_encrypt"

	"github.com/wuqifei/chat/protomodel"

	"github.com/wuqifei/server_lib/libio"
)

var conn net.Conn

func ReadMsg(netConn net.Conn) {
	conn = netConn
	go Ping(conn)

	for {
		lbuf := make([]byte, 2)
		_, err := io.ReadFull(conn, lbuf)
		if err != nil {
			continue
		}

		length := libio.GetUint16BE(lbuf)
		tbuf := make([]byte, length)
		_, err = io.ReadFull(conn, tbuf)
		if err != nil {
			continue
		}
		p := libio.GetUint16BE(tbuf[:2])

		tbuf = tbuf[2:]

		switch p {
		case protomodel.ProtoConfirm:
			{
				ConfirmS2C(conn, tbuf)
			}
		case protomodel.ProtoSendSinglePublishTextMsg:
			{
				RecvMsg(conn, tbuf)
			}
		}

	}
}

func SendMsg(str string) {
	displaySend(str)
	encrtMsg, e := rsa_encrypt.RSAEncrypt([]byte(str), []byte(myInfo.PublicKey))
	if e != nil {
		panic(e)
	}

	content := base64_encrypt.Base64StdEncode(encrtMsg)

	usersJson, err := libfile.ReadfromFile("user.json")
	if err != nil {
		panic(err)
	}

	users := make([]*Info, 0)
	err = json.Unmarshal(usersJson, &users)
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		sendMsgReq := new(protomodel.SendSingleTextMsgC2S)
		sendMsgReq.Uid = uint64(user.ClientId)
		sendMsgReq.Message = content
		conn.Write(packdata(protomodel.ProtoSendSingleTextMsg, sendMsgReq))
	}

}
