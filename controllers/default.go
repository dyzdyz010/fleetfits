package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/fleetfits/models"
)

type MainController struct {
	beego.Controller
}

type Item struct {
	typeID string

	Name struct {
		En string `yaml:"en"`
		Zh string `yaml:"zh"`
	}
}

type ItemList []Item

func (c *MainController) Get() {

	// 	fitStr := `[Loki, Loki]
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
	// fmt.Println(models.ParseFit(fitStr))

	// fmt.Println(models.TypeIDByName("Command Processor I"))
	c.TplNames = "index.tpl"
}

func (c *MainController) Post() {
	fitStr := c.Ctx.Request.FormValue("fitStr")
	fit := models.ParseFit(fitStr)
	fitLink := models.LinkOfFit(fit)
	fmt.Println(fitLink)
	c.Data["FitStr"] = fitStr
	c.Data["FitLink"] = fitLink
	c.Data["FitName"] = fit.Name
	c.Data["ShipName"] = fit.Ship.Name
	c.TplNames = "index.tpl"
}
