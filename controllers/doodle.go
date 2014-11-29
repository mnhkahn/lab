package controllers

import (
	"github.com/astaxie/beego"
	"lab/models"
	"net/http"
)

type DoodleController struct {
	beego.Controller
}

func (this *DoodleController) Get() {
	http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, models.GetDoodle().Doodle, http.StatusFound)
}
