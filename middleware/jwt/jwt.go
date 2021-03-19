package jwt

import (
	"github.com/gin-gonic/gin"
	"time"

	"VueElementAdminGoBackend/models"
	"VueElementAdminGoBackend/pkg/e"
	"VueElementAdminGoBackend/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := util.Gin{C: c}

		var code int

		code = e.Success
		token := c.GetHeader("Access-Token") // todo
		if token == "" {
			code = e.WarningNotLogin
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorCheckAuthTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.WarningAuthTokenTimeout
			} else if claims != nil {
				isExist := models.CheckUser(claims.Username, claims.Password) // 验证用户名密码
				if isExist {                                                  // 用户名、密码验证通过
					c.Set("username", claims.Username)
				} else {
					code = e.WarningAuthAlreadyChange
				}

			}
		}

		if code != e.Success {
			appG.Response(code, nil)

			c.Abort()
			return
		}

		c.Next()
	}
}
