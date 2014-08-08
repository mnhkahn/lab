package controllers

import (
	// "encoding/json"
	// "fmt"
	// "github.com/Max-Liu/Useragent"
	"github.com/astaxie/beego"
	// "io/ioutil"
	// "lab/models"
	// "strings"
)

type IpController struct {
	beego.Controller
}

func (this *IpController) GetIp() {
	// Ip := models.Ip{}

	// c := appengine.NewContext(this.Controller.Ctx.Request)
	// client := urlfetch.Client(c)
	// resp, _ := client.Get("http://61.4.185.48:81/g/")
	// if resp != nil {
	// 	contents, _ := ioutil.ReadAll(resp.Body)

	// 	results := strings.Split(string(contents), ";")
	// 	Ip.Ip = results[0][strings.Index(results[0], "_")+1 : strings.LastIndex(results[0], "_")]
	// 	Ip.Id = results[1][strings.Index(results[1], "=")+1:]
	// }

	// // Ip.Device = this.Controller.Ctx.Input.Header("User-Agent")
	// ua := useragent.NewUserAgent()
	// ua.SetUseragent(this.Controller.Ctx.Request.UserAgent())

	// Ip.UserAgent = ua

	// if len(this.GetString("jsoncallback")) > 0 {
	// 	results, _ := json.Marshal(Ip)
	// 	fmt.Fprintf(this.Ctx.ResponseWriter, this.Controller.Ctx.Request.URL.Query().Get("jsoncallback")+"(["+string(results)+"])")
	// } else {
	// 	this.Data["json"] = &Ip
	// 	this.ServeJson()
	// }

}
