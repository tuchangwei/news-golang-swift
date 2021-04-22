package db

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"server/utils/result"
	"testing"
)

func TestPost_CRUD(t *testing.T) {
	user := InsertUser(t)
	post := &Post{
		Title:    "Hello",
		PostType: 1,
		Content:  "Hello world",
		UserID:   user.ID,
	}
	code, _ := post.Insert()
	assert.Equal(t, code, result.Success)

	code = post.CheckExistViaID()
	assert.Equal(t, code, result.Success)

	post.Title = "Hello1"
	post.PostType = 2
	post.Content = "Hello world 1"

	code, _ = post.Edit()
	assert.Equal(t, post.Title, "Hello1")
	assert.Equal(t, post.PostType, 2)
	assert.Equal(t, post.Content, "Hello world 1")

	p := Post{}
	p.ID = post.ID
	code, _ = p.GetViaID()
	assert.Equal(t, code, result.Success)
	assert.Equal(t, p.Title, post.Title)
	assert.Equal(t, p.PostType, post.PostType)
	assert.Equal(t, p.Content, post.Content)
	assert.Equal(t, p.UserID, post.UserID)

	code, _ = p.DeleteViaID()
	assert.Equal(t, code, result.Success)

	for i := 0; i < 21; i++ {
		post := &Post{
			Title:    fmt.Sprintf("Hello%d", i),
			PostType: 1,
			Content:  "Hello world",
			UserID:   user.ID,
		}
		post.Insert()
	}

	pp := Post{}
	code, _, posts, total := pp.GetAllPosts(int(user.ID), 20, 1)
	assert.Equal(t, code, result.Success)
	assert.Equal(t, len(posts), 1)
	assert.Equal(t, total, int64(21))
}
