package controllers

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"github.com/astaxie/beegae"
	"io/ioutil"
	"lab/models"
	"strings"
)

type WeatherController struct {
	beegae.Controller
}

func (this *WeatherController) GetWeather() {
	Ip := models.Ip{}

	c := appengine.NewContext(this.Controller.Ctx.Request)
	client := urlfetch.Client(c)
	resp, _ := client.Get("http://61.4.185.48:81/g/")
	if resp != nil {
		contents, _ := ioutil.ReadAll(resp.Body)

		results := strings.Split(string(contents), ";")
		Ip.Ip = results[0][strings.Index(results[0], "_")+1 : strings.LastIndex(results[0], "_")]
		Ip.Id = results[1][strings.Index(results[1], "=")+1:]
	}

	Weather := models.Weather{}
	PM := models.PM{}

	resp_weather, _ := client.Get("http://m.weather.com.cn/data/" + Ip.Id + ".html")
	if resp_weather != nil {
		contents, _ := ioutil.ReadAll(resp_weather.Body)
		json.Unmarshal(contents, &Weather)

		resp_pm, _ := client.Get("http://www.pm25.in/api/querys/pm2_5.json?city=" + Weather.Weatherinfo.CityEn + "&token=5j1znBVAsnSf5xQyNQyq")
		contents_pm, _ := ioutil.ReadAll(resp_pm.Body)
		json.Unmarshal(contents_pm, &PM)
		Weather.Weatherinfo.PM2_5 = PM.PM2_5
		Weather.Weatherinfo.Area = PM.Area
		Weather.Weatherinfo.Pm2_5 = PM.Pm2_5
		Weather.Weatherinfo.Pm2_5_24h = PM.Pm2_5_24h
		Weather.Weatherinfo.PositionName = PM.PositionName
		Weather.Weatherinfo.PrimaryPollutant = PM.PrimaryPollutant
		Weather.Weatherinfo.Quality = PM.Quality
		Weather.Weatherinfo.StationCode = PM.StationCode
		Weather.Weatherinfo.TimePoint = PM.TimePoint
	}

	this.Data["json"] = &Weather
	this.ServeJson()
}
