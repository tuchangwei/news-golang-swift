package db

import (
	"server/model"
	"server/utils/result"
)

type PostRepo struct {

}

func NewPostRepo() *PostRepo {
	return &PostRepo{}
}
func (pr *PostRepo) CheckExistViaID(id int) (code int, post model.Post) {
	var p = model.Post{}
	p.ID = uint(id)
	err := DB.Select("*").First(&p).Error
	if err != nil {
		return result.PostNotExist, p
	}
	return result.Success, p
}

func (pr *PostRepo) Insert(post model.Post) (code int, message *string) {
	err := DB.Create(&post).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (pr *PostRepo) DeleteVia(postID int) (code int, message *string) {
	post := model.Post{}
	post.ID = uint(postID)
	err := DB.Delete(post).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}

func (pr *PostRepo) Edit(post model.Post) (code int, message *string) {
	err := DB.Model(&post).Select("title", "content", "post_type").Updates(&post).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg
	}
	return result.Success, nil
}


func (pr *PostRepo) GetVia(postID int) (code int, message *string, post model.Post) {
	var p = model.Post {}
	p.ID = uint(postID)
	err := DB.Model(&p).First(&p).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg, p
	}
	return result.Success, nil, p
}
func (pr *PostRepo) GetAllPosts(userID int, pageSize int, pageNumber int) (code int, message *string, posts []model.Post, count int64) {
	var ps []model.Post
	var query = DB.Model(&model.Post{})
	if userID != 0 {
		query = query.Where("user_id=?", userID)
	}
	var total int64
	query.Count(&total)
	err := DB.Model(&model.Post{}).Limit(pageSize).Offset(pageNumber).Order("created_at desc").Find(&ps).Error
	if err != nil {
		msg := err.Error()
		return result.Error, &msg, ps, total
	}
	return result.Success, nil, ps, total
}
func (pr *PostRepo) DeleteAll() {
	DB.Exec("DELETE FROM posts")
}