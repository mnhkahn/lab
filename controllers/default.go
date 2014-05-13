package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beegae"
)

const JS string = "js"

type CyeamController struct {
	beegae.Controller
}

func (this *CyeamController) Get() {
	this.TplNames = "index.html"
}

func (this *CyeamController) Serve(result interface{}) {
	if len(this.GetString("jsoncallback")) > 0 {
		results, _ := json.Marshal(result)
		fmt.Fprintf(this.Ctx.ResponseWriter, this.Controller.Ctx.Request.URL.Query().Get("jsoncallback")+"(["+string(results)+"])")
	} else {
		this.Data["json"] = &result
		this.ServeJson()
	}
}
