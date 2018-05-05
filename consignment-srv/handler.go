package main

import (
	"context"
	"log"

	pb "github.com/codehakase/go-microservice-example/consignment-srv/proto/consignment"
	vesselProto "github.com/codehakase/go-microservice-example/vessel-srv/proto/vessel"
	"gopkg.in/mgo.v2"
)

type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRespository{s.session.Clone()}
}

// CreateConsignment takes a context and a request, and creates a new consignment.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	// call a client instance of the vessel service with consignment data
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		log.Printf("could not find consignment: %v", err)
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	// set vesselID
	req.VesselId = vesselResponse.Vessel.Id

	// save consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments retrieves all consignments persisted
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
