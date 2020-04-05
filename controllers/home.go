package controllers

import (
	"book/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	if cates, err := new(models.Category).GetCates(-1, 1); err == nil {
		c.Data["Cates"] = cates
	} else {

	}
	c.TplName = "home/list.html"
}
