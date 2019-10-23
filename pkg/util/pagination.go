package util

import (
	"VueElementAdminGoBackend/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) (int, int, int) {
	offset := 0

	page, _ := com.StrTo(c.DefaultQuery("page", "0")).Int()

	limit, _ := com.StrTo(c.DefaultQuery("limit", setting.DefaultLimit)).Int()

	if page > 0 {
		offset = (page - 1) * limit
	} else {
		page = 1
	}

	return offset, page, limit
}
