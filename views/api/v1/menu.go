package v1

import (
	"github.com/DeluxeYang/VueElementAdminGoBackend/pkg/e"
	"github.com/DeluxeYang/VueElementAdminGoBackend/pkg/util"
	"github.com/DeluxeYang/VueElementAdminGoBackend/service/menu_service"
	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	appG := util.Gin{C: c}

	menu, err := menu_service.GetMenu()
	if err != nil {
		appG.Response(e.ErrorGetMenu, nil)
		return
	}

	appG.Response(e.Success, menu)
}

func SaveMenu(c *gin.Context) {
	appG := util.Gin{C: c}

	var menuVOList []menu_service.MenuVO

	err := c.ShouldBind(&menuVOList)
	if err != nil {
		appG.Response(e.WarningInvalidParams, nil)
		return
	}

	err = menu_service.SaveMenu(menuVOList)
	if err != nil {
		appG.Response(e.ErrorSaveMenu, nil)
		return
	}

	appG.Response(e.Success, nil)
}
