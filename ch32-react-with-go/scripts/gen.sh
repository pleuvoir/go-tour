


echo "็ๆ backend ไปฃ็ "

OUT=../pb
protoc \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
common.proto request_service.proto push_service.proto


