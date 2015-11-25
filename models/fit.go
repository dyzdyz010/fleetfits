package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ShipFit struct {
	Name  string    `bson:"name" json:"name"`
	Ship  EVEType   `bson:"ship" json:"ship"`
	Items []EVEType `bson:"items,omitempty" json:"items"`
}

func ParseFit(fitStr string) ShipFit {
	fit := ShipFit{}
	itemArr := strings.Split(fitStr, "\r\n")
	// fmt.Println(len(itemArr))
	shipStr := strings.Replace(itemArr[0], "[", "", -1)
	shipStr = strings.Replace(shipStr, "]", "", -1)
	// fmt.Println(shipStr)

	shiptmp := strings.Split(shipStr, ",")
	shipName := shiptmp[0]
	fitName := strings.Replace(shiptmp[1], " ", "", -1)
	ship := EVEType{}
	err := c_eve_types.Find(bson.M{"name": shipName}).One(&ship)
	if err != nil {
		panic(err)
	}
	fmt.Println(ship)
	fit.Ship = ship
	fit.Name = fitName

	for _, line := range itemArr[1:] {
		fmt.Println(line)
		if line == "" {
			continue
		}

		equipment := EVEType{}
		err := c_eve_types.Find(bson.M{"name": line}).One(&equipment)
		if err != nil {
			if err == mgo.ErrNotFound {
				fit.Items = append(fit.Items, EVEType{Name: mgo.ErrNotFound.Error()})
			} else {
				fmt.Println("%s not found in database.", line)
				panic(err)
			}
		}
		fmt.Println(equipment)
		fit.Items = append(fit.Items, equipment)
	}

	fmt.Println(fit.Items)

	return fit
}

func LinkOfFit(fit ShipFit) string {
	result := fit.Ship.TypeID + ":"

	for _, item := range fit.Items {
		result = result + item.TypeID + ";1" + ":"
	}

	result = result + ":"

	return result
}
