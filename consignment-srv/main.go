package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/codehakase/go-microservice-example/consignment-srv/proto/consignment"
	vesselProto "github.com/codehakase/go-microservice-example/vessel-srv/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	// db conn
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("could not connect to datastore with host %s: %+v", host, err)
	}
	defer session.Close()

	// Create a new service
	srv := micro.NewService(
		// Service name must match package name given in protobuf definition
		micro.Name("go.micro.srv.consignement"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()

	// register our service handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
