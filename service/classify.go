package service

import (
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/model/database"
	"github.com/Demonx24/Vblog-backend/model/other"
	"github.com/Demonx24/Vblog-backend/model/request"
	"github.com/Demonx24/Vblog-backend/utils"
)

type ClassifyService struct {
}

func (classifyservice *ClassifyService) ClassifyList(classify request.ClassifyList) (interface{}, int64, error) {
	db := global.DB
	if classify.ClassId != 0 {
		db = db.Where("class_id=?", classify.ClassId)
	}
	option := other.MySQLOption{
		PageInfo: classify.PageInfo,
		Where:    db,
	}
	option.Order = "class_id desc"
	return utils.MySQLPagination(&database.Classify{}, option)
}
