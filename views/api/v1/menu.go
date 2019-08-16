package v1

import (
	"github.com/DeluxeYang/VueElementAdminGoBackend/pkg/e"
	"github.com/DeluxeYang/VueElementAdminGoBackend/pkg/util"
	"github.com/DeluxeYang/VueElementAdminGoBackend/service/menu_service"
	"github.com/gin-gonic/gin"
)

// @Summary Get menu tree data
// @Produce  json
// @Tags menu
// @Success 200 {object} util.Response "{"code":200,"data":[{"id":1, "menu_name":"", "menu_type":0, "remark":"", "component":"", "permission_tag":"", "path":"", "icon":"", "children":[]}],"message":"ok"}"
// @Router /api/v1/menus [get]
func GetMenu(c *gin.Context) {
	appG := util.Gin{C: c}

	menu, err := menu_service.GetMenu()
	if err != nil {
		appG.Response(e.ErrorGetMenu, nil)
		return
	}

	appG.Response(e.Success, menu)
}

// @Summary Save all menu
// @Produce  json
// @Tags menu
// @Param menu_name body string true "menu_name"
// @Param menu_type body int true "menu_type"
// @Param remark body string true "remark"
// @Param component body string true "component"
// @Param permission_tag body string true "permission_tag"
// @Param path body string true "path"
// @Param icon body string true "icon"
// @Param children body array true "children"
// @Success 200 {object} util.Response "{"code":200,"data":{},"message":"ok"}"
// @Router /api/v1/menus [post]
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
