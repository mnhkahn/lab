package models

import (
	"cyeam_post/dao"
	cyeam_post_model "cyeam_post/models"
	// "fmt"
	"cyeam_post/cygo"
	"html/template"
	"strings"
)

type Post struct {
	Id          interface{}   `json:"id"`
	Title       string        `json:"title`
	CreateTime  string        `json:"create_time"`
	Author      string        `json:"author"`
	Detail      template.HTML `json:"detail"`
	Category    string        `json:"category"`
	Tags        string        `json:"tags"`
	Figure      string        `json:"figure"`
	Description template.HTML `json:"description"`
	Link        string        `json:"link"`
	Source      string        `json:"source"`
	ParseDate   string        `json:"parse_date"`
}

func GetPost(author, sort string, page, size int) []Post {
	Dao, _ := dao.NewDao("solr", "http://128.199.131.129:8983/solr/post")
	models := Dao.GetPost(author, sort, size, (page-1)*size)
	return FormatPost(models, "")
}

func SearchPost(q string, page, size int) (int, float64, []Post) {
	Dao, _ := dao.NewDao("solr", "http://128.199.131.129:8983/solr/post")
	Dao.Debug(true)
	numfound, qtime, models := Dao.Search(q, size, (page-1)*size)
	return numfound, qtime, FormatPost(models, q)
}

func FormatPost(models []cyeam_post_model.Post, key string) []Post {
	posts := make([]Post, 0)
	for _, model := range models {
		post := Post{}
		post.Id = model.Id
		post.Title = model.Title
		post.CreateTime = model.CreateTime.Time.Format(cygo.DATE_LAYOUT_CHINA)
		post.Author = model.Author
		post.Category = model.Category
		post.Tags = model.Tags
		post.Figure = model.Figure
		post.Detail = template.HTML(model.Detail)
		post.Description = template.HTML(model.Description)
		post.Link = model.Link
		post.Source = model.Source
		post.ParseDate = model.ParseDate.Time.Format(cygo.DATE_LAYOUT_CHINA)
		posts = append(posts, post)
	}
	return posts
}

func getDesc(detail, key string) string {
	i := strings.Index(detail, key)
	start := 0
	if i-180 > start {
		start = i - 180
	}
	end := len(detail)
	if i+180 < end {
		end = i + 180
	}
	return string([]byte(detail)[start:end])
}
