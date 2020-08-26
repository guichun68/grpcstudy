## 场景练习(2) POST提交主子订单、嵌套模型

* 启动服务端
```bash
cd server
go run server.go
# switch another terminal
go run httpserver.go
```
* POST查询（或使用postman请求）
```bash
curl -H "Content-Type: application/json" -X POST -d '{"order_id":34,"order_no":"bj00123456","user_id":8901,"order_money":34.12,"order_details":[{"detail_id":101,"order_no":"bj30987","prod_id":789},{"detail_id":102,"order_no":"bj30988","prod_id":790}]}' "http://localhost:8080/v1/orders"
```
响应：
```bash
{
    "status": "OK",
    "message": "success"
}
```
同时server控制台将输出订单详情。