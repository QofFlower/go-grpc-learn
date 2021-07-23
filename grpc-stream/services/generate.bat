@REM cd pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
cd pbfiles && protoc --go_out=../services --go-grpc_out=../services Prod.proto
protoc --go_out=../services --go-grpc_out=../services Models.proto
protoc --go_out=../services --go-grpc_out=../services Order.proto
protoc --go_out=../services --go-grpc_out=../services User.proto

protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Order.proto