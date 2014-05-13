package controllers

import (
	// "github.com/astaxie/beegae"
	"fmt"
	"time"
)

const Layout = "2006-01-02"

var Xianxing = [5][2]int{[2]int{5, 0}, [2]int{1, 6}, [2]int{2, 7}, [2]int{3, 8}, [2]int{4, 9}}

var StartDate = time.Date(2014, 4, 14, 0, 0, 0, 0, time.UTC)

var GAP = 13 * 7 * 24 * time.Hour

type CarController struct {
	CyeamController
}

func (this *CarController) Get() {
	date, err := time.Parse(Layout, this.GetString("date"))
	if err != nil {
		date = time.Now()
	}

	Gap := date.Sub(StartDate).Hours() / 24
	Weeks := int(Gap+1) / 7
	Left := int(Gap+1) % 7

	this.Data["Today0"] = Xianxing[Weeks/13+Left-1][0]
	this.Data["Today1"] = Xianxing[Weeks/13+Left-1][1]
	this.Data["Tomorrow0"] = Xianxing[Weeks/13+Left][0]
	this.Data["Tomorrow1"] = Xianxing[Weeks/13+Left][1]

	Weekday := date.Weekday()
	this.Data["Monday0"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)-3][0]
	this.Data["Monday1"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)-3][1]
	this.Data["Tuesday0"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)-2][0]
	this.Data["Tuesday1"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)-2][1]
	this.Data["Wednesday0"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)-1][0]
	this.Data["Wednesday1"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)-1][1]
	this.Data["Thursday0"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)][0]
	this.Data["Thursday1"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)][1]
	this.Data["Friday0"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)+1][0]
	this.Data["Friday1"] = Xianxing[Weeks/13+Left+int(Weekday-time.Monday)+1][1]

	this.Data["Gap"] = fmt.Sprintf("%d %d", Weeks, Left)
	this.Data["Date"] = date.Format(Layout)
	this.TplNames = "xianxing.html"
}
