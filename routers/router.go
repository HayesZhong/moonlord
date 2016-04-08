package routers

import (
	"github.com/astaxie/beego"
	"moonlord/controllers"
)

func init() {
	beego.Router("/", &controllers.ViewController{}, "*:Index")
	beego.Router("/index", &controllers.ViewController{}, "*:Index")
	beego.Router("/nerest", &controllers.ViewController{}, "*:Nerest")

	beego.Router("/api/getonetras", &controllers.ApiController{}, "*:GetOneTras")
	beego.Router("/api/getnearestmtras", &controllers.ApiController{}, "*:GetNearestMTras")

}
