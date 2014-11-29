package controllers

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"lab/models"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WeixinController struct {
	beego.Controller
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
var WeatherLayout = "【%s】 %s，%s，%s度"

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
		doodle := models.GetDoodle()

		if doodle.Doodle != "" {
			resp.MsgType = News
			resp.Content = "doodle"
			resp.ArticleCount = 1

			a := models.Item{}
			a.Title = doodle.Title
			a.PicUrl = doodle.Doodle
			a.Description = "点击『查看原文』来查看详细说明"
			a.Url = "http://cyeam.com/"
			resp.FuncFlag = 1
			resp.Articles = append(resp.Articles, &a)
		} else {
			resp.Content = fmt.Sprintf("%v", err)
		}
	} else if req.Content == "bing" {
		bing := models.GetBing()
		resp.MsgType = News
		resp.Content = "bing"
		resp.ArticleCount = 1

		a := models.Item{}
		a.Title = "bing"
		a.PicUrl = bing
		a.Description = "点击『查看原文』来查看详细说明"
		a.Url = "http://cyeam.com/"
		resp.FuncFlag = 1
		resp.Articles = append(resp.Articles, &a)
	}

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
