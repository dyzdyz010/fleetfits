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

var types []EVEType

func init() {
	session, err := mgo.Dial("192.168.99.100:32768")
	if err != nil {
		panic(err)
	}
	db = session.DB("evefleetfits")
	c_eve_types = db.C("eve_types")

	// initTypes()

	// fitStr := `[Loki, Loki]
	// Co-Processor II
	// Co-Processor II
	// Co-Processor II
	// Co-Processor II

	// Command Processor I
	// Command Processor I
	// Command Processor I
	// Conjunctive Ladar ECCM Scanning Array I
	// Conjunctive Ladar ECCM Scanning Array I

	// Covert Ops Cloaking Device II
	// Skirmish Warfare Link - Evasive Maneuvers II
	// Skirmish Warfare Link - Rapid Deployment II
	// Skirmish Warfare Link - Interdiction Maneuvers II
	// Information Warfare Link - Sensor Integrity II

	// Medium Processor Overclocking Unit I
	// Medium Processor Overclocking Unit I
	// Medium Low Friction Nozzle Joints I

	// Loki Electronics - Dissolution Sequencer
	// Loki Defensive - Warfare Processor
	// Loki Engineering - Supplemental Coolant Injector
	// Loki Offensive - Covert Reconfiguration
	// Loki Propulsion - Interdiction Nullifier`

	// components := strings.Split(fitStr, "\n")
	// shiptmp := strings.Split(components[0], ",")
	// shiptmp = strings.Split(shiptmp[0], "[")
	// ship := shiptmp[1]
	// fmt.Println(ship)
}

func initTypes() {

	data, err := ioutil.ReadFile("models/types.json")
	if err != nil {
		panic(err)
	}
	types = make([]EVEType, 0)
	json.Unmarshal(data, &types)

	for i := 0; i < len(types); i++ {
		T := types[i]
		c_eve_types.Insert(T)
	}
}
