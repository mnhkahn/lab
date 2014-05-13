package controllers

import (
	"appengine"
	"appengine/urlfetch"
	// "encoding/xml"
	"github.com/astaxie/beegae"
	"io/ioutil"
	"lab/models"
	"regexp"
)

type DoodleController struct {
	beegae.Controller
}

func (this *DoodleController) Get() {
	c := appengine.NewContext(this.Controller.Ctx.Request)
	client := urlfetch.Client(c)
	resp, _ := client.Get("http://www.google.com")
	contents, _ := ioutil.ReadAll(resp.Body)

	cy := ParseDoodle(contents)

	this.Data["json"] = &cy
	this.ServeJson()
}

func ParseDoodle(contents []byte) models.CyeamDoodle {

	// v := models.Rss{}
	// xml.Unmarshal(contents, &v)

	re_img := regexp.MustCompile("<img.*src=(.*?)[^>]*?>")
	// img := re_img.FindAllString(v.Channel.Items[0].Description, -1)
	img := re_img.FindAllString(string(contents), -1)

	re_src := regexp.MustCompile("src=\"?(.*?)(\"|>|\\s+)")
	src := re_src.FindString(img[0])

	url := "http://www.google.com.hk" + src[5:len(src)-1]

	re_title := regexp.MustCompile("alt=\"?(.*?)(\"|>|\\s+)")
	title := re_title.FindString(img[0])
	title += ""

	cy := models.CyeamDoodle{}
	// cy.Title = title[6:8]
	cy.Doodle = url

	return cy
}
