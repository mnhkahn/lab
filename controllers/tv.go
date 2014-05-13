package controllers

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"io/ioutil"
	"lab/models"
	"regexp"
	"strings"
	"time"
)

type TVController struct {
	CyeamController
}

var NBA_LOGO = "http://cyeam.qiniudn.com/%s.jpg"
var NBA = map[string]string{"湖人": "1", "凯尔特人": "2", "热火": "3", "篮网": "4", "尼克斯": "5", "魔术": "6", "76人": "7", "奇才": "8", "活塞": "9", "步行者": "10", "鹈鹕": "11", "雄鹿": "12", "老鹰": "13", "公牛": "14", "猛龙": "15", "骑士": "16", "小牛": "17", "马刺": "18", "森林狼": "19", "爵士": "20", "火箭": "21", "灰熊": "22", "掘金": "23", "国王": "24", "开拓者": "25", "太阳": "26", "勇士": "27", "雷霆": "28", "快船": "29", "山猫": "30"}

func (this *TVController) Get() {
	TVs := []models.TV{}

	c := appengine.NewContext(this.Controller.Ctx.Request)
	client := urlfetch.Client(c)

	now := time.Now()
	resp, _ := client.Get("http://match.sports.sina.com.cn/tvguide/program/top/?date=" + now.Format("2006-01-02"))
	if resp == nil {

	}
	contents, _ := ioutil.ReadAll(resp.Body)
	results := string(contents)

	re_mtypec02 := regexp.MustCompile("<li data-mtype=\"mtypec02\">((.|\\s)+?)</li>")
	mtypec02 := re_mtypec02.FindAllString(results, -1)

	for i := 0; i < len(mtypec02); i++ {
		TV := models.TV{}
		TV.Time = mtypec02[i]

		re_time := regexp.MustCompile("<div class=\"mth_status (mth_status_todo|mth_status_ing|mth_status_done)\">(.*?)</div>")
		time_status := re_time.FindAllString(mtypec02[i], -1)[0]

		if strings.Index(time_status, "<br />") == -1 {
			TV.Time = time_status[strings.Index(time_status, ">")+1 : strings.Index(time_status, "</div")]
			TV.Status = "未开始"
		} else {
			TV.Time = time_status[strings.Index(time_status, ">")+1 : strings.Index(time_status, "<br")]
			TV.Status = time_status[strings.Index(time_status, "/>")+2 : strings.Index(time_status, "</div>")]
		}

		re_title := regexp.MustCompile("<div class=\"mth_title mth_title_font\">(.*?)</div>")
		title := re_title.FindAllString(mtypec02[i], -1)[0]
		team_score := strings.Split(title[strings.Index(title, "<br />")+6:strings.Index(title, "</div>")], " ")

		Team0 := models.Team{}
		Team0.Name = team_score[0]
		Team0.Logo = fmt.Sprintf(NBA_LOGO, GetNBALogo(Team0.Name))
		Team1 := models.Team{}
		Team1.Name = team_score[2]
		Team1.Logo = fmt.Sprintf(NBA_LOGO, GetNBALogo(Team1.Name))

		if strings.Index(team_score[1], "-") == -1 {
			Team0.Score = "0"
			Team1.Score = "0"
		} else {
			Team0.Score = strings.Split(team_score[1], "-")[0]
			Team1.Score = strings.Split(team_score[1], "-")[1]
		}

		TV.Teams = append(TV.Teams, Team0)
		TV.Teams = append(TV.Teams, Team1)

		re_tv := regexp.MustCompile("<img.*src=(.*?)[^>]*?>")
		tvs := re_tv.FindAllString(mtypec02[i], -1)

		for j := 0; j < len(tvs); j++ {
			TV.TVs = append(TV.TVs, tvs[j][strings.Index(tvs[j], "title=")+7:strings.Index(tvs[j], "/>")-2])
		}

		// TV.Text = mtypec02[i]

		TVs = append(TVs, TV)
	}

	this.Serve(TVs)
}

func GetNBALogo(team string) string {
	return NBA[team]
}
