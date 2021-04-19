package model

import "gorm.io/gorm"
type Post struct {
	gorm.Model
	Title    string `json:"title"`
	PostType int    `json:"post_type"` //1 text, 2, image, 3, url
	Content  string `json:"content"`
	UserID uint `json:"user_id"`
}

