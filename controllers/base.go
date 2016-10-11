package controllers

import (
	"github.com/astaxie/beego"
)

var (
	// AccessToken is globally unique for interface call, expires in 7200s, needs to update every 2 hours.
	AccessToken = ""
)

type baseController struct {
	beego.Controller
}

func (base *baseController) Get() {
	base.Data["Website"] = "beego.me"
	base.Data["Email"] = "astaxie@gmail.com"
	base.TplName = "index.tpl"
}
