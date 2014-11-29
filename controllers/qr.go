package controllers

import (
	"github.com/mnhkahn/qrgo"
)

func (this *CyeamController) Qr() {
	code, _ := qr.Encode(this.GetString("url"), qr.H)
	this.Ctx.ResponseWriter.Write(code.PNG())
}
