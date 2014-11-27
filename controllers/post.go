package controllers

import (
	"lab/models"
)

func (this *CyeamController) GetPost() {
	author := this.GetString("author")
	sort := this.GetString("sort")
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	size, _ := this.GetInt("size")
	if size <= 0 {
		size = 1
	}
	view := models.GetPost(author, sort, page, size)
	this.Serve(view)
}
