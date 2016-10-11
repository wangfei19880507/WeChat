package routers

import (
	"github.com/astaxie/beego"

	"weChat/controllers"
	"weChat/controllers/dialogService/basisSupporting"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/dialog/basis/token", &basisSupporting.BasisController{}, "*:GetAccessToken")
	beego.Router("/dialog/basis/ip", &basisSupporting.BasisController{}, "*:GetServerIP")
}
