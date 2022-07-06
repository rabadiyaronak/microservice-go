.PHONY: protos

protos:
	 protoc -I protos/ protos/currency.proto --go_out=protos  --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=protos


	 