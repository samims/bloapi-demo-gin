package models

import "github.com/astaxie/beego/orm"

// Article model
type Article struct {
	ID          int64   `json:"id" orm:"id"`
	Title       *string `json:"title" orm:"column(title)"`
	Description *string `json:"description" orm:"column(description)"`
}

func init() {
	orm.RegisterModel(new(Article))
}
