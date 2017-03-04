package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Redirect("/app/tracking", 302)
}

func (c *MainController) TestQr() {
	c.TplName = "testqr.tpl"
}
