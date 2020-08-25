## 生成证书
* 下载openssl，安装并把目录加入环境变量
* 进入bin（如已加入系统环境变量则忽略本步骤)
* cmd --> openssl
* genrsa -des3 -out server.key 2048 (会生成server.key私钥文件)
* 创建证书请求文件(生成server.csr)： 
```bash
  req -new -key server.key -out server.csr
  #其中Common Name客户端会用到
```
* 删除密码
```shell script
rsa -in server.key -out server_no_passwd.key
```
* 生成服务端证书
```shell script
x509 -req -days 365 -in server.csr -signkey server_no_passwd.key -out server.crt
```
* 拷贝server.crt和server_no_passwd.key到server/keys目录

## 运行
```bash
#启动服务端
$ cd server
$ go run server.go
#启动客户端
$ cd client
$ go run main.go

```