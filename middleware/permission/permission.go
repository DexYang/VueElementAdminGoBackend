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

var Perms = map[string]int{"Retrieve": 2, "Create": 3, "Update": 4, "Delete": 5, "Export": 6}

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
