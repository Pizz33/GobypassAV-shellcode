# GobypassAV-shellcode

cobaltstrike免杀 使用go写的shellcode免杀加载器，实测过 bypass火绒、360、360核晶、360杀毒、def、金山毒霸等主流杀软

2023/4/27 更新了加密方式，提升免杀性

详细教程请移步博客：https://pizz33.github.io/posts/4ac17cb886a9/

食用方法：

1、生成c的payload

![image](https://user-images.githubusercontent.com/88339946/232708666-a8e28b1b-2502-4bbc-91a9-d88e5ff44e9d.png)

2、`go run encode.go` 对shellcode进行加密

3、加密后的结果填到代码里编译运行 `go build -ldflags="-w" decode.go`

![image](https://user-images.githubusercontent.com/88339946/232708833-9709b6c6-59b3-455a-aaa5-e4a92e549c3b.png)


免杀效果：

![image](https://user-images.githubusercontent.com/88339946/232709153-3e45970a-a0ae-409b-bfdc-f9db0209d0ce.png)

![image](https://user-images.githubusercontent.com/88339946/234928250-bcf2952f-c345-4241-b33c-73e053b54dd5.png)

![image](https://user-images.githubusercontent.com/88339946/233016193-23d034da-951a-400a-9720-fffa2b21ba81.png)

![image](https://user-images.githubusercontent.com/88339946/234165227-7a26383c-6f8f-484a-8bfb-6d35d2880e59.png)

![image](https://user-images.githubusercontent.com/88339946/234788023-2a9fd53a-2c02-4467-9ef1-6c654106680d.png)

![image](https://user-images.githubusercontent.com/88339946/232708290-e8f5c3cb-52cb-45bf-a7ea-43615bae0e9d.png)

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
