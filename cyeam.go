package Cyeam

import (
	"github.com/astaxie/beegae"
	"html/template"
	"lab/controllers"
	"net/http"
)

func init() {
	beegae.Errorhandler("404", page_not_found)
	beegae.Router("/ip", &controllers.IpController{}, "*:GetIp")
	beegae.Router("/weather", &controllers.WeatherController{}, "*:GetWeather")
	beegae.Router("/weixin/verify", &controllers.WeixinController{})
	beegae.Router("/weixin", &controllers.WeixinController{})
	beegae.Router("/tv", &controllers.TVController{})
	beegae.Router("/doodle", &controllers.DoodleController{})
	beegae.Router("/bing", &controllers.BingController{})
	beegae.Router("/car", &controllers.CarController{})
	beegae.Run()
}

var errtol = "<iframe scrolling='no' frameborder='0' src='http://yibo.iyiyun.com/js/yibo404/key/2354' width='640' height='464' style='display:block;'></iframe>"

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("beegoerrortemp").Parse(errtol)
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
