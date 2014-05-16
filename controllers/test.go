package controllers

import (
	"time"
)

type TestController struct {
	CyeamController
}

const RFC1123 = "Mon, 02 Jan 2006 15:04:05 CST"

func (this *TestController) Time() {
	this.Serve(GetShanghaiTime())
}

func GetShanghaiTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}
