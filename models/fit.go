package models

import (
// "fmt"
)

type ShipFit struct {
	Name     string   `bson:"name" json:"name"`
	Ship     string   `bson:"ship" json:"ship"`
	Items    []string `bson:"items,omitempty" json:"itmes"`
	Tags     []string `bson:"tags,omitempty" json:"tags"`
	Priority int      `bson:"priority" json:"priority"`
}
