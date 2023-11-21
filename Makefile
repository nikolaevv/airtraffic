protoc:
	protoc internal/service/grpc/pb/airtraffic.proto --go_out=./ --go-grpc_out=./

lint:
	golangci-lint run --timeout 5m

run:
	docker-compose up --build
