# GobypassAV-shellcode
2023/04/18 使用go写的shellcode免杀加载器，免杀主流杀软，bypass火绒、360、def等

食用方法：

1、生成c的payload

![image](https://user-images.githubusercontent.com/88339946/232708666-a8e28b1b-2502-4bbc-91a9-d88e5ff44e9d.png)

2、加密后的结果填到代码里编译运行 `go build -ldflags="-s -w"`

![image](https://user-images.githubusercontent.com/88339946/232708833-9709b6c6-59b3-455a-aaa5-e4a92e549c3b.png)


免杀效果：

![image](https://user-images.githubusercontent.com/88339946/232709153-3e45970a-a0ae-409b-bfdc-f9db0209d0ce.png)

![image](https://user-images.githubusercontent.com/88339946/232708175-24667b5d-09c5-4c43-b8d5-ebb193b17335.png)

![image](https://user-images.githubusercontent.com/88339946/232708290-e8f5c3cb-52cb-45bf-a7ea-43615bae0e9d.png)

![image](https://user-images.githubusercontent.com/88339946/232708368-37c6bd82-8a56-4b15-a298-4576a95fd5ee.png)
