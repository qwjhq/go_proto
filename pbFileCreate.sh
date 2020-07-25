#生成Prod.pb.go
cd pbfiles && protoc --go_out=plugins=grpc:../services/ Prod.proto
#生成网关文件Prod.pb.gw.go
protoc --grpc-gateway_out=logtostderr=true:../services/ Prod.proto
