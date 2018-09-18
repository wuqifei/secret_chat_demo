package cpacket

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/wuqifei/server_lib/libencrypt/aes_ecb_encrypt"

	"github.com/wuqifei/server_lib/libencrypt/base64_encrypt"
	"github.com/wuqifei/server_lib/libencrypt/rsa_encrypt"

	"github.com/wuqifei/server_lib/libfile"
	"github.com/wuqifei/server_lib/libio"

	"github.com/gogo/protobuf/proto"
	"github.com/wuqifei/chat/protomodel"
)

type Info struct {
	ClientId   int64  `json:"client_id"`
	RandomKey  string `json:"random_key"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

var myInfo *Info

func init() {
	myInfo = new(Info)

	content, err := libfile.ReadfromFile("app_public_key.pem")
	if err != nil {
		panic(err)
	}
	myInfo.PublicKey = string(content)

	content, err = libfile.ReadfromFile("app_private_key.pem")
	if err != nil {
		panic(err)
	}
	myInfo.PrivateKey = string(content)

}

func ConfirmS2C(conn net.Conn, val []byte) {
	m := new(protomodel.ConfirmS2C)
	proto.Unmarshal(val, m)
	ret := new(protomodel.ConfirmC2S)
	conn.Write(packdata(protomodel.ProtoConfirm, ret))

	Rows[0] = fmt.Sprintf(Rows[0], m.ClientId)
	Rows[1] = fmt.Sprintf(Rows[1], m.RandomKey)
	myInfo.ClientId = m.ClientId
	myInfo.RandomKey = m.RandomKey

	b, _ := json.Marshal(myInfo)
	libfile.SaveToFile("./me.json", b, nil)
	reoladUI()
}

func RecvMsg(conn net.Conn, val []byte) {
	m := new(protomodel.SendSinglePublishTextMsgS2C)
	proto.Unmarshal(val, m)
	ret := new(protomodel.SendSinglePublishTextMsgC2S)
	conn.Write(packdata(protomodel.ProtoSendSinglePublishTextMsg, ret))
	base64DecodeVal, _ := base64_encrypt.Base64StdDecode(m.Message)
	contentfile, err := libfile.ReadfromFile("./user.json")
	if err != nil {
		panic(err)
	}

	users := make([]*Info, 0)
	err = json.Unmarshal(contentfile, &users)
	if err != nil {
		panic(err)
	}

	var decryptMsg []byte
	for _, user := range users {
		if user.ClientId == int64(m.Uid) {
			decryptVal1 := aes_ecb_encrypt.Decrypt(base64DecodeVal, user.RandomKey)
			decryptMsg, err = rsa_encrypt.RSADecrypt(decryptVal1, []byte(user.PrivateKey))
			if err != nil {
				return
			}
		}
	}
	if decryptMsg == nil {
		return
	}
	displayRecv(m.Uid, string(decryptMsg))
}

func packdata(pr uint16, val proto.Message) []byte {
	valB, _ := proto.Marshal(val)
	p := make([]byte, 2)
	libio.PutUint16BE(p, pr)
	p = append(p, valB...)
	l := uint16(len(p))
	lenB := make([]byte, 2)
	libio.PutUint16BE(lenB, l)
	return append(lenB, p...)
}

// Ping ping pong
func Ping(conn net.Conn) {
	ticker := time.NewTicker(time.Duration(40) * time.Second)
	msg := new(protomodel.PingC2S)

	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			{
				msg.TimeStamp = time.Now().Unix()
				conn.Write(packdata(protomodel.ProtoPing, msg))
			}
		}
	}

}
