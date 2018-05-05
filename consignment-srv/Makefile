build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/codehakase/go-microservice-example/consignment-srv proto/consignment/consignment.proto

	GOOS=linux GOARH=amd64 go build
	# docker build -t consignment-srv .

run:
	docker run -p 50051:50051 consignment-srv \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-srv

run-dev:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/codehakase/go-microservice-example/consignment-srv proto/consignment/consignment.proto

	MICRO_SERVER_ADDRESS=:50051 MICRO_REGISTRY=mdns go run main.go
