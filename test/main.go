package main

import (
	"fmt"
)

type Info struct {
	ClientId   int64  `json:"client_id"`
	RandomKey  string `json:"random_key"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

func main() {
	// str := "test123"

	// myInfo := new(Info)
	// content, _ := libfile.ReadfromFile("/home/wqf/Documents/go_work_space/work/src/github.com/wuqifei/chat/client/me.json")
	// json.Unmarshal(content, myInfo)

	// rsaencrypt, _ := rsa_encrypt.RSAEncrypt([]byte(str), []byte(myInfo.PublicKey))
	// base64EncStr := base64_encrypt.Base64StdEncode(rsaencrypt)

	// base64Dec, _ := base64_encrypt.Base64StdDecode(base64EncStr)

	// aesEnc := aes_ecb_encrypt.Encrypt(base64Dec, myInfo.RandomKey)
	// base64EncStr = base64_encrypt.Base64StdEncode(aesEnc)

	// base64Dec, _ = base64_encrypt.Base64StdDecode(base64EncStr)
	// aesDec := aes_ecb_encrypt.Decrypt(base64Dec, myInfo.RandomKey)
	// rsaDescyp, _ := rsa_encrypt.RSADecrypt(aesDec, []byte(myInfo.PrivateKey))
	// fmt.Printf("[%s]\n", string(rsaDescyp))

	for i := 0; i < 64; i++ {
		ssql := `update video_` + fmt.Sprintf("%d", i) + ` set width = '640' where width = '480' and height = '360' and media_key  like 'production/stchat/video/%';`
		fmt.Printf("%s\n", ssql)

	}

}
