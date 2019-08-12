package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/DeluxeYang/VueElementAdminGoBackend/models"
	"github.com/DeluxeYang/VueElementAdminGoBackend/pkg/e"
	"github.com/DeluxeYang/VueElementAdminGoBackend/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username" binding:"required"`
	Password string `valid:"Required; MaxSize(50)" json:"password" binding:"required"`
}

func GetAuth(c *gin.Context) {
	appG := util.Gin{C: c}

	authJson := auth{}
	code := e.WarningInvalidParams

	data := make(map[string]interface{})

	if c.BindJSON(&authJson) == nil { // 绑定json
		username := authJson.Username // 获取用户名
		password := authJson.Password

		valid := validation.Validation{} // 验证
		ok, _ := valid.Valid(&authJson)  // 验证数据完整性

		if ok { // 数据完整性验证成功

			isExist := models.CheckUser(username, password) // 验证用户名密码
			if isExist {                                    // 如果存在用户名、密码

				token, err := util.GenerateToken(username, password) // 生成JWT token

				if err != nil {
					code = e.ErrorAuthTokenGenerate // token生成错误
				} else {
					data["token"] = token
					code = e.Success
				}

			} else {
				code = e.WarningWrongAuth
			}

		} else { // 数据完整性验证失败
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message) // 打印日志
			}
		}
	}

	appG.Response(code, data)
}
