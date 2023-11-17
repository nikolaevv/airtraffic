protoc:
	protoc internal/service/grpc/pb/airtraffic.proto --go_out=./ --go-grpc_out=./
