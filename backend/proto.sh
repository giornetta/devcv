protoc \
    -I proto \
    -I vendor/ \
    --go_out=plugins=grpc:proto \
    --grpc-gateway_out=logtostderr=true:proto \
    proto/developers.proto