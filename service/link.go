package service

import (
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/model/database"
)

type LinkService struct {
}

func (linkService *LinkService) GetAllLinks() ([]database.Link, error) {
	db := global.DB
	var links []database.Link
	err := db.Find(&links).Error
	return links, err
}
