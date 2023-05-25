redeployed-at:=$(shell date +%s)

.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/google/wire/cmd/wire@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golang/mock/mockgen@latest
	@echo "Installing protoc-gen-validate (PGV) can currently only be done from source. See: https://github.com/envoyproxy/protoc-gen-validate#installation"

.PHONY: update
update:
	go get -u ./...
	go mod tidy

.PHONY: buf
buf:
	cd api && buf mod update
	cd api && buf generate

.PHONY: wire
wire:
	wire ./...

.PHONY: mock
mock:
	mockgen -source=./api/gen/go/user/v1/userservice_grpc.pb.go -destination=./mock/mock_userservice_grpc.pb.go -package=mock
	mockgen -source=./api/gen/go/order/v1/orderservice_grpc.pb.go -destination=./mock/mock_orderservice_grpc.pb.go -package=mock

.PHONY: test
test:
	go test -cover -race -covermode=atomic -coverprofile=coverage.txt ./...


.PHONY: user-server
user-server:
	go run ./src/user/cmd/main.go

.PHONY: order-server
order-server:
	go run ./src/order/cmd/main.go


