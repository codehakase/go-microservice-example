package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/codehakase/go-microservice-example/vessel-srv/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("could not connect to datastore %s,: %v", host, err)
	}
	defer session.Close()

	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register implementation
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}
