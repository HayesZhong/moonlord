package routers

import (
	"moonlord/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ViewController{}, "*:Index")
	beego.Router("/index", &controllers.ViewController{}, "*:Index")
	beego.Router("/pointnerest", &controllers.ViewController{}, "*:PointNerest")

	beego.Router("/api/nearestmtras", &controllers.ApiController{}, "*:GetNearestMTrasByPoint")
	beego.Router("/api/trasnum", &controllers.ApiController{}, "*:GetTrasNum")
	beego.Router("/api/getsimtra", &controllers.ApiController{}, "*:GetSimMTrasByTras")
}
