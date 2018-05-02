package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/codehakase/go-microservice-example/consignment-srv/proto/consignment"
	vesselProto "github.com/codehakase/go-microservice-example/vessel-srv/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	port = ":50051"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository here simulates the use of a datastore.
type ConsignmentRespository struct {
	consignments []*pb.Consignment
}

// Service should implement all the methods to satisfy the service defined in the protobuf definition.
type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

func (repo *ConsignmentRespository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *ConsignmentRespository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// CreateConsignment takes a context and a request, and creates a new consignment.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// call a client instance of the vessel service with consignment data
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	// set vesselID
	req.VesselId = vesselResponse.Vessel.Id

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
	repo := &ConsignmentRespository{}

	// Create a new service
	srv := micro.NewService(
		// Service name must match package name given in protobuf definition
		micro.Name("go.micro.srv.consignement"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()

	// register our service handler
	pb.RegisterShippingServiceServer(srv.Server(), &service{repo, vesselClient})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
