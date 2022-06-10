CURRENT_DIR=$(shell pwd)
proto-gen:
	protoc --proto_path=${CURRENT_DIR} --go-grpc_out=${CURRENT_DIR} --go-grpc_opt=require_unimplemented_servers=false --go_out=${CURRENT_DIR} ${CURRENT_DIR}/*_protos/*.proto ${CURRENT_DIR}/*_protos/*.proto
