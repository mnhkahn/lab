package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"lab/models"
	"regexp"
	"time"
)

type DoodleController struct {
	beego.Controller
}

func (this *DoodleController) Get() {
	req := httplib.Get("http://www.google.com/doodles/doodles.xml")
	req.SetTimeout(time.Duration(5)*time.Second, time.Duration(5)*time.Second)
	req.Debug(beego.AppConfig.String("runmode") == "dev")
	contents, err := req.Bytes()

	if err != nil {
		fmt.Println(err)
		this.Abort("500")
	}
	cy := ParseDoodle(contents)

	this.Data["json"] = &cy
	this.ServeJson()
}

func ParseDoodle(contents []byte) models.CyeamDoodle {
	v := models.Rss{}
	xml.Unmarshal(contents, &v)

	re_img := regexp.MustCompile("<img.*src=(.*?)[^>]*?>")
	img := re_img.FindAllString(v.Channel.Items[0].Description, -1)

	re_src := regexp.MustCompile("src=\"?(.*?)(\"|>|\\s+)")
	src := re_src.FindString(img[0])

	url := "http:" + src[5:len(src)-1]

	cy := models.CyeamDoodle{}
	cy.Title = v.Channel.Items[0].Title
	cy.Doodle = url

	return cy
}
