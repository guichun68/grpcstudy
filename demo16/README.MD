## 流模式入门（上）
场景：批量查询用户积分



* 启动服务端
```bash
cd server
go run server.go
# switch another terminal
go run httpserver.go
```
* POST正常查询：
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
