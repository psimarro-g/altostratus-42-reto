package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Asteroid struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name          string             `json:"name,omitempty" validate:"required"`
	Diameter      float64            `json:"diameter,omitempty" validate:"required"`
	DiscoveryDate string             `json:"discovery_date,omitempty" validate:"required"`
	Observations  string             `json:"observations"`
	//Distances	  Distances[] 	`json:"distances"`
}
