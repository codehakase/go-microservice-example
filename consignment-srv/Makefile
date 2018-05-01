build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/codehakase/go-microservice-example/consignment-srv proto/consignment/consignment.proto

	GOOS=linux GOARH=amd64 go build
	docker build -t consignment-srv .

run:
	docker run -p 50051:50051 consignment-srv
