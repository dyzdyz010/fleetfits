package models

import (
	"encoding/json"
	// "fmt"
	"gopkg.in/mgo.v2"
	"os"
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
	if _, err := os.Stat("models/types.json"); err == nil {
		data, err := ioutil.ReadFile("models/types.json")
		if err != nil {
			panic(err)
		}
		types = make([]EVEType, 0)
		json.Unmarshal(data, &types)

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

		err = os.Rename("models/types.json", "models/types.json.backup")
		if err != nil {
			panic(err)
		}
	}
}
