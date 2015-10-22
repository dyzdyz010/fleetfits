package routers

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/fleetfits/controllers"
)

func init() {
	// Pages
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{}, "get:Dashboard")
	beego.Router("/admin/fit/:id", &controllers.AdminController{}, "get:Fit")

	// APIs
	beego.Router("/admin/fit/parse", &controllers.AdminController{}, "post:ParseFit")
}
