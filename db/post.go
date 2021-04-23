package db

import (
	"gorm.io/gorm"
	"server/utils/result"
)
type Post struct {
	gorm.Model
	Title    string `json:"title"`
	PostType int    `json:"post_type" gorm:"not null;default:1"` //1 text, 2, image, 3, url
	Content  string `json:"content"`
	UserID uint `json:"user_id"`
}
func (p *Post) CheckExistViaID() (code int) {
	err := DB.Select("*").First(&p).Error
	if err != nil {
		return result.PostNotExist
	}
	return result.Success
}

func (p *Post) Insert() (code int, message *string) {
	err := DB.Create(&p).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (p *Post) DeleteViaID() (code int, message *string) {
	err := DB.Delete(p, p.ID).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (p *Post) Edit() (code int, message *string) {
	err := DB.Model(p).Select("title", "content", "post_type").Updates(p).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}


func (p *Post) GetViaID() (code int, message *string) {
	err := DB.Model(p).Where("id=?", p.ID).First(&p).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}
func (p *Post) GetAllPosts(userID int, pageSize int, pageNumber int) (code int, message *string, posts []Post, count int64) {
	var ps []Post
	var query = DB.Model(&Post{})
	if userID != 0 {
		query = query.Where("user_id=?", userID)
	}
	var total int64
	query.Count(&total)
	if pageSize == 0 {
		pageSize = 20
	}
	err := query.Limit(pageSize).Offset(pageNumber*pageSize).Order("created_at desc").Find(&ps).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg, ps, total
	}
	return result.Success, nil, ps, total
}
