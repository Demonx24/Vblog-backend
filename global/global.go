package global

import (
	"github.com/Demonx24/Vblog-backend/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
