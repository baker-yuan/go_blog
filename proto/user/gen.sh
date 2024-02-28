echo "生成 rpc 代码"

# 输出目录
OUT=./

# grpc代码生成
protoc \
-I ${OUT} \
-I $GOPATH/src \
-I $GOPATH/src/google/api/*.proto \
-I $GOPATH/src/validate/*.proto \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
--validate_out="lang=go:./" \
--go-gin_out=${OUT} \
user.proto



protoc \
-I ${OUT} \
-I $GOPATH/src \
-I $GOPATH/src/google/api/*.proto \
-I $GOPATH/src/validate/*.proto \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
--validate_out="lang=go:./" \
--grpc-gateway_out=${OUT} \
--grpc-gateway_opt=paths=source_relative \
user.proto



