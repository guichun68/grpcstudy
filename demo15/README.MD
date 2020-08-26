## 穿插知识点 使用第三方库进行字段验证
介绍一个第三方社区库  
https://github.com/envoyproxy/protoc-gen-validate  
安装： 
```bash
go get -u github.com/envoyproxy/protoc-gen-validate
```
安装好后$GOPATH/bin下会生成 protoc-gen-validate.exe文件



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
* 当价格不符合规则时的查询：
```bash
curl -H "Content-Type: application/json" -X POST -d '{"order_id":34,"order_no":"bj00123456","user_id":8901,"order_money":1,"order_details":[{"detail_id":101,"order_no":"bj30987","prod_id":789},{"detail_id":102,"order_no":"bj30988","prod_id":790}]}' "http://localhost:8080/v1/orders"
```
响应：
```bash
{
    "status": "error",
    "message": "invalid OrderMain.OrderMoney: value must be greater than 1"
}
```
