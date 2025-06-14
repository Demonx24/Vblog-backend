package service

import (
	"errors"
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/model/database"
	"github.com/Demonx24/Vblog-backend/model/other"
	"github.com/Demonx24/Vblog-backend/model/request"
	"github.com/Demonx24/Vblog-backend/utils"
	"gorm.io/gorm"
	"sort"
	"time"
)

type BlogService struct {
}

func (blogService *BlogService) BloggerId(blog database.Blog) (database.Blog, error) {
	db := global.DB
	err := db.Where("id = ?", blog.Id).First(&blog).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return database.Blog{}, errors.New("文章不存在")
		}
		return database.Blog{}, errors.New("查询id失败")
	}
	return blog, nil
}
func (blogService *BlogService) CreateBlog(blog database.Blog) (database.Blog, error) {
	db := global.DB
	blog.Createdate = time.Now()
	blog.Status = 1
	blog.Lang = "cn"
	if err := db.Create(&blog).Error; err != nil {
		return database.Blog{}, err
	}
	return blog, nil
}
func (blogService *BlogService) DeleteBlog(blog database.Blog) error {
	db := global.DB
	if err := db.Where("id=? ", blog.Id).Find(&blog).Error; err != nil {
		return errors.New("查询失败")
	}
	if blog.Status == 0 {
		return errors.New("文章以被删除")
	}
	if err := db.Model(&blog).Where("id=?", blog.Id).Update("status", 0).Error; err != nil {
		return err
	}
	return nil
}
func (blogService *BlogService) UpdateBlog(blog database.Blog) error {
	db := global.DB
	blog.Createdate = time.Now()
	if err := db.Model(&blog).Where("id=?", blog.Id).Updates(blog).Error; err != nil {
		return err
	}
	return nil
}
func (blogService *BlogService) BlogList(blogs request.BlogList) (interface{}, int64, error) {
	db := global.DB
	if blogs.BlogId != 0 {
		db = db.Where("blog_id=?", blogs.BlogId)
	}
	option := other.MySQLOption{
		PageInfo: blogs.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.Blog{}, option)
}
func (blogService *BlogService) GetCard1(card1 database.Card1) (database.Card1, error) {
	db := global.DB

	// 查询中文文章总数
	err := db.Raw("SELECT SUM(CASE WHEN lang='cn' THEN 1 ELSE 0 END) AS cnTotal FROM blog").Scan(&card1.CnTotal).Error
	if err != nil {
		return card1, err
	}

	// 查询分类总数
	err = db.Raw("SELECT COUNT(*) AS cateTotal FROM classify").Scan(&card1.CateTotal).Error
	if err != nil {
		return card1, err
	}

	// 查询标签总数
	err = db.Raw("SELECT COUNT(*) AS tagTotal FROM tag").Scan(&card1.TagTotal).Error
	if err != nil {
		return card1, err
	}

	return card1, nil
}
func (blogService *BlogService) GetCard4() ([]database.Blog, error) {
	db := global.DB

	var blogs []database.Blog
	err := db.Where("status = ? AND lang = ?", 1, "cn").
		Order("createdate DESC").
		Limit(5).
		Find(&blogs).Error
	if err != nil {
		return []database.Blog{}, err
	}

	return blogs, nil
}
func (blogService *BlogService) GetArchive() (*other.ArchiveResult, error) {
	var blogs []database.Blog
	db := global.DB
	err := db.Where("lang = ?", "cn").
		Order("createdate ASC").
		Find(&blogs).Error
	if err != nil {
		return nil, err
	}

	blogGroups := make(map[string][]database.Blog)
	for _, blog := range blogs {
		yearMonth := blog.Createdate.Format("2006-01")
		blogGroups[yearMonth] = append(blogGroups[yearMonth], blog)
	}

	for _, group := range blogGroups {
		sort.Slice(group, func(i, j int) bool {
			return group[i].Createdate.After(group[j].Createdate)
		})
	}

	var total int64
	err = db.Model(&database.Blog{}).Where("lang = ?", "cn").Count(&total).Error
	if err != nil {
		return nil, err
	}

	return &other.ArchiveResult{
		BlogGroups: blogGroups,
		Total:      total,
	}, nil
}
func (blogService *BlogService) GetIndex() ([]database.Blog, error) {
	var blogs []database.Blog
	db := global.DB
	err := db.
		Where("status = ? AND lang = ?", 1, "cn").
		Order("createdate desc").
		Limit(10).
		Find(&blogs).Error

	return blogs, err
}
