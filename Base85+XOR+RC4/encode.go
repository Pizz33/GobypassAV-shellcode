package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"

	"github.com/eknkc/basex"
)

func main() {
	key := []byte("demaxiya")                                         // RC4 加密使用的密钥
	message := "\xfc\x48\x83\x41\x01\xc1\x56\x6d\x00\x05\xf5\xe1\x00" // 原始消息

	// XOR 操作
	xordMessage := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		xordMessage[i] = message[i] ^ 0xff
	}

	// RC4 加密
	cipher, _ := rc4.NewCipher(key)
	rc4Message := make([]byte, len(xordMessage))
	cipher.XORKeyStream(rc4Message, xordMessage)

	// 转为十六进制
	hexCiphertext := make([]byte, hex.EncodedLen(len(rc4Message)))
	n := hex.Encode(hexCiphertext, rc4Message)
	hexCiphertext = hexCiphertext[:n]

	// Base85 编码
	base85, _ := basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	encodedMessage := base85.Encode(hexCiphertext)

	fmt.Println(encodedMessage)
}
