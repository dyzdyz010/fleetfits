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

	fmt.Println(models.TypeIDByName("Command Processor I"))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
