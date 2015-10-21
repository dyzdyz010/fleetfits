package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/dyzdyz010/fleetfits/models"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Dashboard() {
	c.TplNames = "admin.tpl"
}

func (c *AdminController) Fit() {

	fitID := c.Ctx.Input.Param(":id")

	c.Data["FitID"] = fitID
	c.TplNames = "fit.tpl"
}
