package controllers

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"github.com/astaxie/beegae"
	"io/ioutil"
	"lab/models"
	// "strings"
)

type WeatherController struct {
	beegae.Controller
}

func (this *WeatherController) GetWeather() {
	c := appengine.NewContext(this.Controller.Ctx.Request)
	client := urlfetch.Client(c)
	resp, _ := client.Get("http://api.map.baidu.com/telematics/v3/weather?location=%E5%8C%97%E4%BA%AC&output=json&ak=43E57D0f47ca6382344892b9b65ba0ad")
	Weather := models.Weather{}
	contents, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(contents, &Weather)

	big := true
	for i := 0; i < len(Weather.Results[0].WeatherDate); i++ {
		Weather.Results[0].WeatherDate[i].PicUrl = GetWeatherIcon(Weather.Results[0].WeatherDate[i].Weather, big)
		if i == 0 {
			big = false
		}
	}

	this.Data["json"] = &Weather
	this.ServeJson()
}

func GetWeatherIcon(Weather string, big bool) string {
	Icon := "http://cyeam.qiniudn.com/1400265638_weather-sunny.png"
	if big {
		if Weather == "晴" {
			if GetShanghaiTime().Hour() > 6 && GetShanghaiTime().Hour() < 18 {
				Icon = "http://cyeam.qiniudn.com/1400265636_weather-sunny.png"
			} else {
				Icon = "http://cyeam.qiniudn.com/1400265527_weather-moon.png"
			}
		} else if Weather == "多云" || Weather == "多云转阴" {
			Icon = "http://cyeam.qiniudn.com/1400265312_weather-partlycloudy.png"
		} else if Weather == "阵雨转晴" {
			Icon = "http://cyeam.qiniudn.com/1400265562_weather-thunder-rainy-h.png"
		} else {
			Icon = "http://cyeam.qiniudn.com/1400265636_weather-sunny.png"
		}
	} else {
		if Weather == "晴" {
			if GetShanghaiTime().Hour() > 6 && GetShanghaiTime().Hour() < 18 {
				Icon = "http://cyeam.qiniudn.com/1400265636_weather-sunny.png"
			} else {
				Icon = "http://cyeam.qiniudn.com/1400265528_weather-moon.png"
			}
		} else if Weather == "多云" || Weather == "多云转阴" {
			Icon = "http://cyeam.qiniudn.com/1400265348_weather-partlycloudy.png"
		} else if Weather == "阵雨转晴" {
			Icon = "http://cyeam.qiniudn.com/1400265564_weather-thunder-rainy-h.png"
		}
	}
	return Icon
}
