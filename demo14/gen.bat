cd server/pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services Models.proto
protoc --go_out=plugins=grpc:../services Orders.proto

protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Orders.proto
cd ../..