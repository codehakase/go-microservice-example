package main

import (
	pb "github.com/codehakase/go-microservice-example/consignment-srv/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shipsrv"
	consignmentCollection = "consignment"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRespository struct {
	session *mgo.Session
}

// Create a new consignment
func (repo *ConsignmentRespository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetAll retreives all consignments
func (repo *ConsignmentRespository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

// Close closes the database session after each query has ran
func (repo *ConsignmentRespository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRespository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}
