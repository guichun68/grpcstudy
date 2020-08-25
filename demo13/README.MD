## 场景练习(1) POST方式提交订单数据
-- gateway实现api/http





* 启动服务端
```bash
cd server
go run server.go
# switch another terminal
go run httpserver.go
```
* POST查询（或使用postman请求）
```bash
curl -H "Content-Type: application/json" -X POST -d '{"order_id":34,"order_no":"bj00123456","user_id":8901,"order_money":34.12}' "http://localhost:8080/v1/orders"
```