package controllers

import (
	"book/models"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user/iidex.html"
}

func (c *MainController) Hi() {
	models.PrintUser()
	models.PrintCompanyByOrm()
}
