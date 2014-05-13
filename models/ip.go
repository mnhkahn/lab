package models

import (
	"github.com/Max-Liu/Useragent"
)

type Ip struct {
	Ip        string               `json:"ip"`
	Id        string               `json:"id"`
	UserAgent *useragent.UserAgent `json:"useragent"`
}
