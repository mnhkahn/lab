package controllers

import (
	"fmt"
	"lab/models"
	"regexp"
)

func (this *CyeamController) Regexp() {
	if this.Ctx.Input.IsPost() {
		form := models.RegForm{}
		if err := this.ParseForm(&form); err != nil {
			panic(err)
		}
		// form.Raw = "paranormal"
		// form.Reg = "a."
		fmt.Println(form.Raw, form.Reg)
		re := regexp.MustCompile(form.Reg)
		form.Res = re.FindAllString(form.Raw, -1)
		fmt.Println(form.Res, len(form.Reg))
		this.Data["Form"] = form
	}
	this.TplNames = "regexp.html"
}
