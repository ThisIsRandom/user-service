.PHONY: proto
proto:
	protoc --go_out=./ --go-grpc_out=./ ./proto/service.proto

build:
	docker build -t user-service -f dev.dockerfile ./