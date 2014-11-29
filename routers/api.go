package routers

import (
	"github.com/astaxie/beego"
	"html/template"
	"lab/controllers"
	"net/http"
)

func init() {
	beego.Errorhandler("404", page_not_found)
	beego.Router("/post", &controllers.CyeamController{}, "*:GetPost")
	beego.Router("/doodle", &controllers.DoodleController{})
	beego.Router("/bing", &controllers.BingController{})
}

var errtol = "<iframe scrolling='no' frameborder='0' src='http://yibo.iyiyun.com/js/yibo404/key/2354' width='640' height='464' style='display:block;'></iframe>"

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("beegoerrortemp").Parse(errtol)
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
