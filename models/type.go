package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EVEType struct {
	TypeID string `bson:"_id" json:"typeID"`
	Name   string `json:"name"`
}

func TypeIDByName(name string) string {
	result := EVEType{}
	err := c_eve_types.Find(bson.M{"name": name}).One(&result)

	if err != nil {
		if err == mgo.ErrNotFound {
			return ""
		} else {
			panic(err)
		}
	}

	return result.TypeID
}
