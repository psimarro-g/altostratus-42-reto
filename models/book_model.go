package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Distances struct {
	Date     string  `json:"date"`
	Distance float64 `json:"distance"`
}

type Asteroid struct {
	ID            primitive.ObjectID `bson:"id" json:"id,omitempty"`
	Name          string             `json:"name,omitempty" validate:"required"`
	Diameter      float64            `json:"diameter,omitempty" validate:"required"`
	DiscoveryDate string             `json:"discovery_date,omitempty" validate:"required"`
	Observations  string             `json:"observations"`
	Distances     []Distances        `json:"distances"`
}
