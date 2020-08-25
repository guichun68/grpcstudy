## 语法速学(4) 日期类型、创建主订单模型


需求：  
 在模型文件(Models.proto)中添加订单模型，添加下单时间属性并使用之。

* 使用时间类型
```shell script
import "google/protobuf/timestamp.proto";
```

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