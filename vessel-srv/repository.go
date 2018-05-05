package main

import (
	pb "github.com/codehakase/go-microservice-example/vessel-srv/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shipsrv"
	vesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable checks a specification against a list of vessels
// if capacity and max weight are below a vessel's capacity and max weight
// return that vessel
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
