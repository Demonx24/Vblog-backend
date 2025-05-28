package service

import (
	"errors"
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/model/database"
	"gorm.io/gorm"
)

type BlogService struct {
}

func (blogService *BlogService) BloggerId(blog database.Blog) (database.Blog, error) {
	db := global.DB
	err := db.Where("id = ?", blog.Id).First(&blog).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return database.Blog{}, errors.New("博客不存在")
		}
		return database.Blog{}, errors.New("查询id失败")
	}
	return blog, nil
}
