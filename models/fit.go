package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"regexp"
	"strconv"
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

	for _, typeName := range itemArr[1:] {
		fmt.Println(typeName)
		if typeName == "" {
			continue
		}

		equipment := EVEType{}
		amount := 1

		typeName = strings.Split(typeName, ",")[0]

		re := regexp.MustCompile(".*x[0-9]*$")
		if re.MatchString(typeName) {
			amountSepIndex := strings.LastIndex(typeName, " ")
			etype := typeName[:amountSepIndex]
			amountStr := strings.Replace(typeName[amountSepIndex:], " ", "", -1)
			amount, _ = strconv.Atoi(strings.Replace(amountStr, "x", "", 1))
			typeName = etype
		}

		err := c_eve_types.Find(bson.M{"name": typeName}).One(&equipment)
		if err != nil {
			if err == mgo.ErrNotFound {
				fit.Items = append(fit.Items, EVEType{Name: mgo.ErrNotFound.Error()})
			} else {
				fmt.Println("%s not found in database.", typeName)
				panic(err)
			}
		}
		equipment.Amount = amount
		fmt.Println(equipment)
		fit.Items = append(fit.Items, equipment)
	}

	// fmt.Println(fit.Items)

	return fit
}

func LinkOfFit(fit ShipFit) string {
	result := fit.Ship.TypeID + ":"

	for _, item := range fit.Items {
		result = result + item.TypeID + ";" + strconv.Itoa(item.Amount) + ":"
	}

	result = result + ":"

	return result
}
