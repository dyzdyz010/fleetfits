package models

import (
	"encoding/json"
	// "fmt"
	"gopkg.in/mgo.v2"
	// "strings"
	// "gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

var db *mgo.Database
var c_eve_types *mgo.Collection
var c_config *mgo.Collection

var types []EVEType

func init() {
	session, err := mgo.Dial("192.168.99.100:32768")
	if err != nil {
		panic(err)
	}
	db = session.DB("evefleetfits")
	c_eve_types = db.C("eve_types")
	c_config = db.C("config")

	initTypes()
}

func initTypes() {

	data, err := ioutil.ReadFile("models/types.json")
	if err != nil {
		panic(err)
	}
	types = make([]EVEType, 0)
	json.Unmarshal(data, &types)

	count, err := c_eve_types.Count()
	if err != nil {
		panic(err)
	}
	// fmt.Println(count)
	if count == len(types) {
		return
	}

	for i := 0; i < len(types); i++ {
		T := types[i]

		buf := EVEType{}
		err = c_eve_types.FindId(T.TypeID).One(&buf)
		if err != nil && err != mgo.ErrNotFound {
			panic(err)
		}
		if buf.Name == T.Name {
			continue
		}

		c_eve_types.Insert(T)
	}
}
