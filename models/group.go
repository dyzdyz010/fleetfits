package models

import (
	"gopkg.in/mgo.v2"
)

type FleetGroup struct {
	Name     string   `bson:"name" json:"name"`
	Tags     []string `bson:"tags,omitempty" json:"tags"`
	Fleets   []string `bson:"fleets,omitempty" json:"fleets"`
	Priority int      `bson:"priority" json:"priority"`
}

func FleetGroupByID(id string) FleetGroup {
	result := FleetGroup{}
	err := c_eve_types.FindId(id).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return result
		} else {
			panic(err)
		}
	}

	return result
}
