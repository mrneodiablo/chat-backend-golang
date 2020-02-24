package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = "{ Maria Ozawa, 404 }"
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = "{ Sori AOI, 500 }"
	c.ServeJSON()
}

func (c *ErrorController) ErrorDb() {
	c.Data["json"] = "{ Nguyen PTK, 503 }"
	c.ServeJSON()
}
