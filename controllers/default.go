package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Home() {
	fmt.Println("test-web")
	this.Data["Website"] = "hahah"
	this.Data["Email"] = "asdfs"
	this.TplName = "index.tpl"
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
