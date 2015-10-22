package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ShipFit struct {
	Name     string   `bson:"name" json:"name"`
	Ship     string   `bson:"ship" json:"ship"`
	Items    []string `bson:"items,omitempty" json:"items"`
	Tags     []string `bson:"tags,omitempty" json:"tags"`
	Priority int      `bson:"priority" json:"priority"`
}

func ParseFit(fitBytes string) ShipFit {
	fit := ShipFit{}
	buf1 := strings.Split(fitBytes, "\n")
	fmt.Println(string(buf1[0]))

	shiptmp := strings.Split(buf1[0], ",")
	shiptmp = strings.Split(shiptmp[0], "[")
	shipName := shiptmp[1]
	ship := EVEType{}
	err := c_eve_types.Find(bson.M{"name": shipName}).One(&ship)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf1[1:])
	fit.Ship = ship.Name

	for _, line := range buf1[1:] {
		fmt.Println(line)
		if line == "" {
			continue
		}

		equipment := EVEType{}
		err := c_eve_types.Find(bson.M{"name": line}).One(&equipment)
		if err != nil {
			if err == mgo.ErrNotFound {
				fit.Items = append(fit.Items, mgo.ErrNotFound.Error())
			} else {
				panic(err)
			}
		}
		fit.Items = append(fit.Items, equipment.Name)
	}

	fmt.Println(fit)

	return fit
}
