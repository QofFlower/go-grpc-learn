@REM cd pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
cd pbfiles && protoc --go_out=../services --go-grpc_out=../services Prod.proto