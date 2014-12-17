package routers

import (
	"github.com/astaxie/beego"
	"lab/controllers"
)

func init() {
	beego.Router("/regexp", &controllers.CyeamController{}, "*:Regexp")
}
