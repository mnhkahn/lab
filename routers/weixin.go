package routers

import (
	"github.com/astaxie/beego"
	"lab/controllers"
)

func init() {
	beego.Router("/weixin/verify", &controllers.WeixinController{})
	beego.Router("/weixin", &controllers.WeixinController{})
}
