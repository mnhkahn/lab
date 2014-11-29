package controllers

import (
	"github.com/astaxie/beego"
	"lab/models"
	"net/http"
)

type BingController struct {
	beego.Controller
}

func (this *BingController) Get() {
	http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, models.GetBing(), http.StatusFound)
}
