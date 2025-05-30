package other

import "github.com/Demonx24/Vblog-backend/model/database"

type ArchiveResult struct {
	BlogGroups map[string][]database.Blog
	Total      int64
}
