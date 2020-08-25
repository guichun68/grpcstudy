## 双向认证下rpc-gateway使用
####（同时提供rpc和http接口） 
第三方库地址：https://github.com/grpc-ecosystem/grpc-gateway

同时提供gRPC和RESTful风格的api。

![Alt text](https://camo.githubusercontent.com/e75a8b46b078a3c1df0ed9966a16c24add9ccb83/68747470733a2f2f646f63732e676f6f676c652e636f6d2f64726177696e67732f642f3132687034435071724e5046686174744c5f63496f4a707446766c41716d35774c513067677149356d6b43672f7075623f773d37343926683d333730 "optional title")

* 修改server/pbfiles/Prod.proto, 添加option相关代码
* 重新生成Prod.pb.go
```bash
cd server/pbfiles
protoc --go_out=plugins=grpc:../services Prod.proto
```
```bash
cd server/pbfiles
protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto

```

运行grpc服务
```bash
cd server
go run server.go
```
运行httpServer服务
```bash
cd server
go run httpserver.go
```
打开浏览器访问
```bash
http://localhost:8080/v1/prod/20
```