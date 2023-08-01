# GobypassAV-shellcode

cobaltstrike免杀，实测过 bypass火绒、360、360核晶、360杀毒、def、金山毒霸等主流杀软

| 杀软类型    | 免杀绕过技巧    |
| ---- | ---- |
| 火绒 | 编译参数限制较多，对hash和字符串特征进行识别，静态能过动态基本不查杀，对很多go库调用报毒。。。|
| 360 | 单360查杀力不高，装了杀毒后直接儿子变爸爸，查杀力大大提升，对于简单的加密识别度较高，容易上线后云查杀过一会掉线，推荐使用分离加载方式，并使用反沙箱的代码延长马子时间|
| 360核晶 | 开启后对整体查杀性能影响不大，避免使用进程注入的方式加载shellcode，无法执行大部分cmd命令和相关程序（使用bof插件进行替代）|
| Defender | 新增许多cobaltstrike规则，推荐使用`Stageless`，免杀性比`Stage`好，4.5版本开启`sleep_mask`参数增强免杀性，对体积大的文件查杀度不高|


详细教程请移步博客：https://pizz33.github.io/posts/4ac17cb886a9/

食用方法：

1、生成c的payload

![image](https://user-images.githubusercontent.com/88339946/232708666-a8e28b1b-2502-4bbc-91a9-d88e5ff44e9d.png)

2、`go run encode.go` or `python xor64.py` 对shellcode进行加密

![image](https://user-images.githubusercontent.com/88339946/232708833-9709b6c6-59b3-455a-aaa5-e4a92e549c3b.png)

3、加密后的结果填到代码里编译运行 `go build decode.go` 

远程加载把加密后的字符串放到云端，把云端地址填到对应位置生成 （可放到vps上或使用oss云存储等）

(这里大多报错为缺少依赖，运行 `go mod init` & `go mod tidy` 拉取即可)

免杀效果：

![image](https://user-images.githubusercontent.com/88339946/234937098-ba1f7e9b-0c8e-4455-a84b-46a6ae53159f.png)

![image](https://user-images.githubusercontent.com/88339946/234936629-b80e9b97-8a85-485e-9097-bbf4091a4d39.png)

![image](https://user-images.githubusercontent.com/88339946/234928250-bcf2952f-c345-4241-b33c-73e053b54dd5.png)

![image](https://user-images.githubusercontent.com/88339946/233016193-23d034da-951a-400a-9720-fffa2b21ba81.png)

![image](https://user-images.githubusercontent.com/88339946/234165227-7a26383c-6f8f-484a-8bfb-6d35d2880e59.png)

![image](https://user-images.githubusercontent.com/88339946/234788023-2a9fd53a-2c02-4467-9ef1-6c654106680d.png)

![image](https://user-images.githubusercontent.com/88339946/232708290-e8f5c3cb-52cb-45bf-a7ea-43615bae0e9d.png)

[![Star History Chart](https://api.star-history.com/svg?repos=Pizz33/GobypassAV-shellcode&type=Date)](https://star-history.com/#star-history/star-history&Date)

项目仅供进行学习研究，切勿用于任何非法未授权的活动，如个人使用违反安全相关法律，后果与本人无关

站在巨人的肩膀上学习，参考借鉴以下师傅的项目，特别感谢

https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-virtualalloc

https://github.com/7BitsTeam/EDR-Bypass-demo 

https://www.yuque.com/aufeng/aufeng_good/aq09p0#yNorm

https://mp.weixin.qq.com/s/xiFbSE6goKFqLAlyACi83A

https://github.com/timwhitez/Doge-Loader

https://github.com/TideSec/GoBypassAV

https://www.crisprx.top/archives/515

https://github.com/Ne0nd0g/go-shellcode

https://github.com/piiperxyz/AniYa

https://github.com/safe6Sec/GolangBypassAV
