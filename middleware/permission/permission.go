package permission

import (
	"VueElementAdminGoBackend/pkg/e"
	"VueElementAdminGoBackend/pkg/util"
	"VueElementAdminGoBackend/service/menu_service"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

const (
	Retrieve = 2
	Create   = 3
	Update   = 4
	Delete   = 5
	Export   = 6
)

type ButtonType struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	Type  int    `json:"type"`
}

var ButtonTypes = [...]ButtonType{
	{Title: "查询", Name: "Retrieve", Type: 2},
	{Title: "新增", Name: "Create", Type: 3},
	{Title: "编辑", Name: "Update", Type: 4},
	{Title: "删除", Name: "Delete", Type: 5},
	{Title: "导出", Name: "Export", Type: 6},
}

func Permission(path string, permission ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := util.Gin{C: c}

		usernameValue, exists := c.Get("username")
		if exists {
			username := com.ToStr(usernameValue)

			pass, err := menu_service.CheckPermission(username, path, permission)
			if err != nil {
				appG.Response(e.ErrorCheckPermission, nil)
				c.Abort()
			} else if !pass {
				appG.Response(e.WarningNoPermission, nil)
				c.Abort()
			}
		}

		c.Next()
	}
}
