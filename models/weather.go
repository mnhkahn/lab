package models

type Weather struct {
	Weatherinfo Weatherinfo `json:"weatherinfo"`
}

type Weatherinfo struct {
	City             string `json:"city"`
	CityEn           string `json:"city_en"`
	DateY            string `json:"date_y"`
	Date             string `json:"date"`
	Week             string `json:"week"`
	Fchh             string `json:"fchh"`
	Cityid           string `json:"cityid"`
	Temp1            string `json:"temp1"`
	Temp2            string `json:"temp2"`
	Temp3            string `json:"temp3"`
	Temp4            string `json:"temp4"`
	Temp5            string `json:"temp5"`
	Temp6            string `json:"temp6"`
	Weather1         string `json:"weather1"`
	Weather2         string `json:"weather2"`
	Weather3         string `json:"weather3"`
	Weather4         string `json:"weather4"`
	Weather5         string `json:"weather5"`
	Weather6         string `json:"weather6"`
	Img1             string `json:"img1"`
	Img2             string `json:"img2"`
	Img3             string `json:"img3"`
	Img4             string `json:"img4"`
	Img5             string `json:"img5"`
	Img6             string `json:"img6"`
	Img7             string `json:"img7"`
	Img8             string `json:"img8"`
	Img9             string `json:"img9"`
	Img10            string `json:"img10"`
	Img11            string `json:"img11"`
	Img12            string `json:"img12"`
	Img_single       string `json:"img_single"`
	Img_title1       string `json:"img_title1"`
	Img_title2       string `json:"img_title2"`
	Img_title3       string `json:"img_title3"`
	Img_title4       string `json:"img_title4"`
	Img_title5       string `json:"img_title5"`
	Img_title6       string `json:"img_title6"`
	Img_title7       string `json:"img_title7"`
	Img_title8       string `json:"img_title8"`
	Img_title9       string `json:"img_title9"`
	Img_title10      string `json:"img_title10"`
	Img_title11      string `json:"img_title11"`
	Img_title12      string `json:"img_title12"`
	Img_title_single string `json:"img_title_single"`
	Wind1            string `json:"wind1"`
	Wind2            string `json:"wind2"`
	Wind3            string `json:"wind3"`
	Wind4            string `json:"wind4"`
	Wind5            string `json:"wind5"`
	Wind6            string `json:"wind6"`
	Fx1              string `json:"fx1"`
	Fx2              string `json:"fx2"`
	Fl1              string `json:"fl1"`
	Fl2              string `json:"fl2"`
	Fl3              string `json:"fl3"`
	Fl4              string `json:"fl4"`
	Fl5              string `json:"fl5"`
	Fl6              string `json:"fl6"`
	Index            string `json:"index"`
	Index_d          string `json:"index_d"`
	Index48          string `json:"index48"`
	Index48_d        string `json:"index48_d"`
	Index_uv         string `json:"index_uv"`
	Index48_uv       string `json:"index48_uv"`
	Index_xc         string `json:"index_xc"`
	Index_tr         string `json:"index_tr"`
	Index_co         string `json:"index_co"`
	Index_cl         string `json:"index_cl"`
	Index_ls         string `json:"index_ls"`
	Index_ag         string `json:"index_ag"`
	St1              string `json:"st1"`
	St2              string `json:"st2"`
	St3              string `json:"st3"`
	St4              string `json:"st4"`
	St5              string `json:"st5"`
	St6              string `json:"st6"`
	PM2_5            string `json:"pm2_5"`
	Area             string `json:"area"`
	Pm2_5            string `json:"pm2_5"`
	Pm2_5_24h        string `json:"pm2_5_24h"`
	PositionName     string `json:"position_name"`
	PrimaryPollutant string `json:"primary_pollutant"`
	Quality          string `json:"quality"`
	StationCode      string `json:"station_code"`
	TimePoint        string `json:"time_point"`
}

type PM struct {
	PM2_5            string `json:"pm2_5"`
	Area             string `json:"area"`
	Pm2_5            string `json:"pm2_5"`
	Pm2_5_24h        string `json:"pm2_5_24h"`
	PositionName     string `json:"position_name"`
	PrimaryPollutant string `json:"primary_pollutant"`
	Quality          string `json:"quality"`
	StationCode      string `json:"station_code"`
	TimePoint        string `json:"time_point"`
}
