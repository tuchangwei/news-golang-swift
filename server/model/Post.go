package model

import "gorm.io/gorm"
type Post struct {
	gorm.Model
	Title    string `json:"title"`
	PostType int    `json:"post_type"`
	Content  string `json:"content"`
	UserID int `json:"_"`
}

