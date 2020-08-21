package routers

import (
	"googleGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/google", &controllers.GoogleController{})
	beego.Router("/callback", &controllers.CallbackController{})
}
