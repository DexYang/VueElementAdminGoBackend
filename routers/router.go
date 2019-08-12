package routers

import (
	"github.com/DeluxeYang/GinProject/middleware/permission"
	"github.com/gin-gonic/gin"

	"github.com/DeluxeYang/GinProject/middleware/jwt"
	"github.com/DeluxeYang/GinProject/pkg/setting"
	"github.com/DeluxeYang/GinProject/views/api"
	"github.com/DeluxeYang/GinProject/views/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.POST("/api/v1/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// 获取用户权限信息
		apiV1.GET("/info", v1.GetUserInfo)

		// 分页获取用户数据
		apiV1.GET("/users", permission.Permission("users", permission.Retrieve), v1.GetUsers)
		// 添加单个用户
		apiV1.POST("/users", permission.Permission("users", permission.Create), v1.AddUser)
		// 获取单个用户
		apiV1.GET("/users/:id", permission.Permission("users", permission.Retrieve), v1.GetUser)
		// 编辑单个用户
		apiV1.PUT("/users/:id", permission.Permission("users", permission.Update), v1.EditUser)
		// 删除单个用户
		apiV1.DELETE("/users/:id", permission.Permission("users", permission.Delete), v1.DeleteUser)

		// 分页获取角色数据
		apiV1.GET("/roles", permission.Permission("roles", permission.Retrieve), v1.GetRoles)
		// 添加单个角色
		apiV1.POST("/roles", permission.Permission("roles", permission.Create), v1.AddRole)
		// 获取单个角色
		apiV1.GET("/roles/:id", permission.Permission("roles", permission.Retrieve), v1.GetRole)
		// 编辑单个角色
		apiV1.PUT("/roles/:id", permission.Permission("roles", permission.Update), v1.EditRole)
		// 删除单个角色
		apiV1.DELETE("/roles/:id", permission.Permission("roles", permission.Delete), v1.DeleteRole)

		// 获取全部目录菜单
		apiV1.GET("/menus", permission.Permission("menus", permission.Retrieve), v1.GetMenu)
		// 保存全部目录菜单
		apiV1.POST("/menus", permission.Permission("menus", permission.Create, permission.Update, permission.Delete), v1.SaveMenu)
	}

	return r
}