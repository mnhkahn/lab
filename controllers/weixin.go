package controllers

import (
	"appengine"
	"appengine/urlfetch"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beegae"
	"io"
	"io/ioutil"
	"lab/models"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WeixinController struct {
	beegae.Controller
}

const (
	TOKEN    = "cyeam"
	Text     = "text"
	Location = "location"
	Image    = "image"
	Link     = "link"
	Event    = "event"
	Music    = "music"
	News     = "news"
)

var TV = "[%s] %s %s : %s %s | %s"
var XianXingTxt = "今日限行尾号：%d和%d"

func (this *WeixinController) Get() {
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")

	dict := []string{timestamp, nonce, echostr}
	sort.Strings(dict)

	h := sha1.New()
	io.WriteString(h, strings.Join(dict, ""))

	if Signature(timestamp, nonce) == signature {
		fmt.Fprintf(this.Ctx.ResponseWriter, echostr)
	} else {
		fmt.Fprintf(this.Ctx.ResponseWriter, "")
	}

}

func (this *WeixinController) Post() {
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}

	var wreq *models.Request
	if wreq, err = DecodeRequest(body); err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}

	wresp, err := dealwith(wreq, this.Ctx.Request)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	data, err := wresp.Encode()
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	this.Ctx.WriteString(string(data))
	return
}

func Signature(timestamp, nonce string) string {
	strs := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func DecodeRequest(data []byte) (req *models.Request, err error) {
	req = &models.Request{}
	if err = xml.Unmarshal(data, req); err != nil {
		return
	}
	req.CreateTime *= time.Second
	return
}

func NewResponse() (resp *models.Response) {
	resp = &models.Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	return
}

func dealwith(req *models.Request, r *http.Request) (resp *models.Response, err error) {
	resp = NewResponse()
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = Text

	resp.Content = req.Content
	if req.Content == "doodle" {
		c := appengine.NewContext(r)
		client := urlfetch.Client(c)
		resp_doodle, _ := client.Get("http://doodle.cyeam.com/")
		contents, _ := ioutil.ReadAll(resp_doodle.Body)

		doodle := models.CyeamDoodle{}
		err := json.Unmarshal(contents, &doodle)

		if err == nil {
			resp.MsgType = News
			resp.Content = "doodle"
			resp.ArticleCount = 1

			a := models.Item{}
			a.Title = doodle.Title
			a.PicUrl = doodle.Doodle
			a.Description = "点击『查看原文』来查看接口"
			a.Url = "http://doodle.cyeam.com/"
			resp.FuncFlag = 1
			resp.Articles = append(resp.Articles, &a)
		} else {
			resp.Content = fmt.Sprintf("%v", err)
		}
	} else if req.Content == "tv" {
		TVs := GetTVs(r)

		resp.MsgType = Text
		resp.Content = ""
		for i := 0; i < len(TVs); i++ {
			resp.Content += fmt.Sprintf(TV, TVs[i].Time, TVs[i].Teams[0].Name, TVs[i].Teams[0].Score, TVs[i].Teams[1].Score, TVs[i].Teams[1].Name, strings.Join(TVs[i].TVs, ",")) + "\n"
		}
	} else if req.Content == "car" {
		Car := GetCar(GetShanghaiTime())

		resp.MsgType = Text
		resp.Content = fmt.Sprintf(XianXingTxt, Car.Today[0], Car.Today[1])
	}
	// if strings.Trim(strings.ToLower(req.Content), " ") == "help" || req.Content == "Hello2BizUser" || req.Content == "subscribe" {
	// 	resp.Content = "目前支持包的使用说明及例子的说明，这些例子和说明来自于github.com/astaxie/gopkg，例如如果你想查询strings有多少函数，你可以发送：strings，你想查询strings.ToLower函数，那么请发送：strings.ToLower"
	// 	return resp, nil
	// }
	// strs := strings.Split(req.Content, ".")
	// var resurl string
	// var a models.Item
	// if len(strs) == 1 {
	// 	resurl = "https://raw.github.com/astaxie/gopkg/master/" + strings.Trim(strings.ToLower(strs[0]), " ") + "/README.md"
	// 	a.Url = "https://github.com/astaxie/gopkg/tree/master/" + strings.Trim(strings.ToLower(strs[0]), " ") + "/README.md"
	// } else {
	// 	var other []string
	// 	for k, v := range strs {
	// 		if k < (len(strs) - 1) {
	// 			other = append(other, strings.Trim(strings.ToLower(v), " "))
	// 		} else {
	// 			other = append(other, strings.Trim(strings.Title(v), " "))
	// 		}
	// 	}
	// 	resurl = "https://raw.github.com/astaxie/gopkg/master/" + strings.Join(other, "/") + ".md"
	// 	a.Url = "https://github.com/astaxie/gopkg/tree/master/" + strings.Join(other, "/") + ".md"
	// }
	// beego.Info(resurl)
	// rsp, err := http.Get(resurl)
	// if err != nil {
	// 	resp.Content = "不存在该包内容"
	// 	return nil, err
	// }
	// defer rsp.Body.Close()
	// if rsp.StatusCode == 404 {
	// 	resp.Content = "找不到你要查询的包:" + req.Content
	// 	return resp, nil
	// }
	// resp.MsgType = News
	// resp.ArticleCount = 1
	// body, err := ioutil.ReadAll(rsp.Body)
	// beego.Info(string(body))
	// a.Description = string(body)
	// a.Title = req.Content
	// a.PicUrl = "http://bbs.gocn.im/static/image/common/logo.png"
	// resp.Articles = append(resp.Articles, &a)
	// resp.FuncFlag = 1

	return resp, nil
}

/*
{
   "button": [
      {
         "name": "Board",
         "sub_button": [
            {
               "type": "click",
               "name": "比赛转播",
               "key": "TV"
            },
            {
               "type": "click",
               "name": "天气",
               "key": "WEATHER"
            },
            {
               "type": "click",
               "name": "限行",
               "key": "CAR"
            }
         ]
      },
      {
         "type": "view",
         "name": "Cyeam",
         "url": "http://www.cyeam.com"
      },
      {
         "type": "view",
         "name": "Blog",
         "url": "http://blog.cyeam.com"
      }
   ]
}
*/
