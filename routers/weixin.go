package routers

import (
	"github.com/astaxie/beegae"
	"lab/controllers"
)

func init() {
	beegae.Router("/weixin/verify", &controllers.WeixinController{})
	beegae.Router("/weixin", &controllers.WeixinController{})
}
