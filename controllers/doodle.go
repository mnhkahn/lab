package controllers

import (
	"encoding/xml"
	"github.com/astaxie/beego"
	// "io/ioutil"
	"lab/models"
	"regexp"
)

type DoodleController struct {
	beego.Controller
}

func (this *DoodleController) Get() {
	// c := appengine.NewContext(this.Controller.Ctx.Request)
	// client := urlfetch.Client(c)
	// resp, _ := client.Get("http://www.google.com/doodles/doodles.xml")
	// contents, _ := ioutil.ReadAll(resp.Body)

	// cy := ParseDoodle(contents)

	// this.Data["json"] = &cy
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
