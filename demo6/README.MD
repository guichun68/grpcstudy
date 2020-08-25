## 使用CA证书  
* 根证书（root certificate）是属于根证书颁发机构（CA)的公钥证书。用以验证它所签发的证书（客户端、服务端）
* 创建ca证书
```bash
$ openssl
> genrsa -out ca.key 2048
# 生成证书请求
> req -new -x509 -days 3650 -key ca.key -out ca.pem
```
* 重新生成服务端证书
```bash
> genrsa -out server.key 2048
# 生成证书请求文件
> req -new -key server.key -out server.csr
# (注意 common name请填写localhost)
# 签发服务器端证书
> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
```
* 生成客户端证书
```bash
> ecparam -genkey -name secp384r1 -out client.key
# 生成证书请求文件
> req -new -key client.key -out client.csr
#签发客户端证书
> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
```

* 拷贝ca.pem、server.key、server.pem到server/cert目录
* 拷贝ca.pem、client.key、client.pem到client/cert目录


* 启动服务端
```bash
cd server
go run server.go
```
* 启动客户端连接并查询
```bash
cd client
go run main.go
```
