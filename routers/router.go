package routers

import (
	"github.com/astaxie/beego"
	"test-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
