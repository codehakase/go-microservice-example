package main

import (
	"context"
	"fmt"

	pb "github.com/codehakase/go-microservice-example/consignment-srv/proto/consignment"
	micro "github.com/micro/go-micro"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository here simulates the use of a datastore.
type Respository struct {
	consignments []*pb.Consignment
}

// Service should implement all the methods to satisfy the service defined in the protobuf definition.
type service struct {
	repo IRepository
}

func (repo *Respository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Respository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// CreateConsignment takes a context and a request, and creates a new consignment.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// save consignment
	consigment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	res.Created = true
	res.Consignment = consigment
	return nil
}

// GetConsignments retrieves all consignments persisted
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignment = consignments
	return nil
}

func main() {
	repo := &Respository{}

	// Create a new service
	srv := micro.NewService(
		// Service name must match package name given in protobuf definition
		micro.Name("go.micro.srv.consignement"),
		micro.Version("latest"),
	)
	srv.Init()

	// register our service handler
	pb.RegisterShippingServiceServer(srv.Server(), &service{repo})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
