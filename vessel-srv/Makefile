build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/codehakase/go-microservice-example/vessel-srv proto/vessel/vessel.proto
    GOOS=linux GOARCH=amd64 go build

		#docker build -t vessel-srv .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=50052 -e MICRO_REGISTRY=mdns vessel-srv
