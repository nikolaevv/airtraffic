protoc:
	protoc internal/service/grpc/pb/airtraffic.proto --go_out=./ --go-grpc_out=./

lint:
	golangci-lint run --timeout 5m

run:
	docker-compose up --build

gen:
	go generate ./internal/...

test:
	CGO_ENABLED=1 go test -p 1 -race -cover -count=1 -coverprofile=.coverprofile -coverpkg=./... ./internal/...
