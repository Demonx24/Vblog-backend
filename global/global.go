package global

import (
	"github.com/Demonx24/Vblog-backend/config"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config     *config.Config
	DB         *gorm.DB
	Log        *zap.Logger
	Redis      redis.Client
	BlackCache local_cache.Cache
)
