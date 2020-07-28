#生成商品pb文件Prod.pb.go
cd pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
#生成订单pb文件
protoc --go_out=plugins=grpc:../services Orders.proto
protoc --go_out=plugins=grpc:../services Users.proto
#生成网关文件Prod.pb.gw.go
protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Orders.proto
# go get -u github.com/envoyproxy/protoc-gen-validate
protoc --go_out="plugins=grpc:../services" --validate_out="lang=go:../services" Models.proto
##  GO111MODULE=off go get -d github.com/envoyproxy/protoc-gen-validate
# cp ../envoyproxy/protoc-gen-validate/validate/* pbfiles/validate/
# cp -r  ../protocolbuffers//protobuf/src/google/protobuf/* pbfiles/protobuf/
# mkdir -p  pbfiles/google/api/
#cp go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis/google/api/* pbfiles/google/api/