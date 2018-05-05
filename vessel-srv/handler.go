package main

import (
	"context"

	pb "github.com/codehakase/go-microservice-example/vessel-srv/proto/vessel"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

// FindAvailable retreives a vessel
func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	// Find the next available vessel
	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}
	// set vessel as part of response
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
