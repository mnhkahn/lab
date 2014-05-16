package controllers

import (
	"lab/models"
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
		date = GetShanghaiTime()
	}

	Gap := date.Sub(StartDate).Hours() / 24
	Weeks := int(Gap+1) / 7
	Left := int(Gap+1) % 7

	Car := models.Car{}

	Car.Today = append(Car.Today, Xianxing[Weeks/13+Left-1][0])
	Car.Today = append(Car.Today, Xianxing[Weeks/13+Left-1][1])
	if true {
		Car.Tomorrow = append(Car.Tomorrow, -1)
		Car.Tomorrow = append(Car.Tomorrow, -1)
	} else {
		Car.Tomorrow = append(Car.Tomorrow, Xianxing[Weeks/13+Left][0])
		Car.Tomorrow = append(Car.Tomorrow, Xianxing[Weeks/13+Left][1])
	}

	this.Serve(Car)
}
