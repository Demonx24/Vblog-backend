package other

import (
	"github.com/Demonx24/Vblog-backend/model/request"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"gorm.io/gorm"
)

type MySQLOption struct {
	request.PageInfo
	Order   string
	Where   *gorm.DB
	Preload []string
}

type EsOption struct {
	request.PageInfo
	Index          string
	Request        *search.Request
	SourceIncludes []string
}
