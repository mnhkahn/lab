package models

import (
	"cyeam_post/dao"
	cyeam_post_model "cyeam_post/models"
	"time"
)

type Post struct {
	Id          int       `json:"id"`
	Title       string    `json:"title`
	CreateTime  time.Time `json:"create_time"`
	Author      string    `json:"author"`
	Detail      string    `json:"detail"`
	Category    string    `json:"category"`
	Tags        string    `json:"tags"`
	Figure      string    `json:"figure"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Source      string    `json:"source"`
	ParseDate   time.Time `json:"parse_date"`
}

func GetPost(author, sort string, page, size int) []Post {
	Dao, _ := dao.NewDao("db", "cyeam:qwerty@tcp(128.199.131.129:3306)/cyeam?charset=utf8")
	models := Dao.GetPost(author, sort, size, (page-1)*size)
	return FormatPost(models)
}

func SearchPost(q string, page, size int) []Post {
	Dao, _ := dao.NewDao("db", "cyeam:qwerty@tcp(128.199.131.129:3306)/cyeam?charset=utf8")
	models := Dao.Search(q, size, (page-1)*size)
	return FormatPost(models)
}

func FormatPost(models []cyeam_post_model.Post) []Post {
	posts := make([]Post, 0)
	for _, model := range models {
		post := Post{}
		post.Id = model.Id
		post.Title = model.Title
		post.CreateTime = model.CreateTime
		post.Author = model.Author
		post.Detail = model.Detail
		post.Category = model.Category
		post.Tags = model.Tags
		post.Figure = model.Figure
		post.Description = model.Description
		post.Link = model.Link
		post.Source = model.Source
		post.ParseDate = model.ParseDate
		posts = append(posts, post)
	}
	return posts
}
