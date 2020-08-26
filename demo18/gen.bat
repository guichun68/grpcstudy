cd server/pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services Orders.proto
protoc --go_out=plugins=grpc:../services --validate_out=lang=go:../services Models.proto

protoc --go_out=plugins=grpc:../services Users.proto

protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Orders.proto
protoc --grpc-gateway_out=logtostderr=true:../services Users.proto

cd ../..