package service

import (
	"errors"
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/model/database"
)

type TagService struct {
}

func (tagService *TagService) TagId(tag database.Tag) (database.Tag, error) {
	db := global.DB
	if err := db.Where("tag_id=?", tag.TagId).Find(&tag).Error; err != nil {
		return database.Tag{}, errors.New("查询失败")
	}
	return tag, nil
}
func (tagService *TagService) AddTag(tag database.Tag) (database.Tag, error) {
	db := global.DB
	if err := db.Create(&tag).Error; err != nil {
		return database.Tag{}, errors.New("<UNK>")
	}
	return tag, nil
}
func (tagService *TagService) DeleteTag(tag database.Tag) error {
	db := global.DB
	if err := db.Where("tag_id=?", tag.TagId).Delete(&database.Tag{}).Error; err != nil {
		return errors.New("删除失败")
	}
	return nil
}
func (tagService *TagService) UpdateTag(tag database.Tag) (database.Tag, error) {
	db := global.DB
	if err := db.Model(&tag).Updates(tag).Error; err != nil {
		return database.Tag{}, errors.New("修改失败")
	}
	return tag, nil
}
func (tagService *TagService) GetAllTags() ([]database.Tag, error) {
	db := global.DB
	var tags []database.Tag
	err := db.Find(&tags).Error
	return tags, err
}
