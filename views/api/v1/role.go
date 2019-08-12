package v1

import (
	"github.com/DeluxeYang/GinProject/pkg/e"
	"github.com/DeluxeYang/GinProject/pkg/util"
	"github.com/DeluxeYang/GinProject/service/menu_service"
	"github.com/DeluxeYang/GinProject/service/role_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	appG := util.Gin{C: c}

	offset, page, limit := util.GetPage(c)
	key := c.DefaultQuery("key", "")

	roles, err := role_service.GetRoleList(offset, limit, key)
	if err != nil {
		appG.Response(e.ErrorGetRoleList, nil)
		return
	}

	total, err := role_service.GetRoleTotal(key)
	if err != nil {
		appG.Response(e.ErrorGetRoleTotal, nil)
		return
	}

	appG.Response(e.Success, map[string]interface{}{
		"list": roles,
		"total": total,
		"page": page,
		"limit": limit,
	})
}

func RoleIDCheck(id int) int {
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		return e.WarningInvalidParams
	}

	exists, err := role_service.ExistRoleByID(id)
	if err != nil {
		return e.ErrorCheckRoleExist
	}
	if !exists {
		return e.WarningRoleNotExist
	}
	return e.Success
}

func GetRole(c *gin.Context)  {
	appG := util.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	if code := RoleIDCheck(id); code != e.Success {
		appG.Response(code, nil)
		return
	}

	user, err := role_service.GetRole(id)
	if err != nil {
		appG.Response(e.ErrorGetRole, nil)
		return
	}

	appG.Response(e.Success, user)
}

func DeleteRole(c *gin.Context) {
	appG := util.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	if code := RoleIDCheck(id); code != e.Success {
		appG.Response(code, nil)
		return
	}

	user, err := role_service.DeleteRole(id)
	if err != nil {
		appG.Response(e.ErrorDeleteRole, nil)
		return
	}

	appG.Response(e.Success, user)
}

func AddRole(c *gin.Context)  {
	appG := util.Gin{C: c}

	roleVO := role_service.RoleVO{}

	code := util.BindAndValid(c, &roleVO) // 数据验证
	if code != e.Success {
		appG.Response(code, nil)
		return
	}

	exists, err := role_service.ExistRoleByRoleName(roleVO.RoleName, 0) //
	if err != nil {
		appG.Response(e.ErrorRoleNameAlreadyExist, nil)
		return
	}
	if exists {
		appG.Response(e.WarningRoleNameAlreadyExist, nil)
		return
	}

	_, err = menu_service.ExistMenuList(roleVO.Menu) //
	if err != nil {
		appG.Response(e.ErrorRoleMenuList, nil)
		return
	}

	resRoleVO, err := role_service.AddRole(&roleVO)
	if err != nil { // 添加User
		appG.Response(e.ErrorAddRole, nil)
		return
	}

	appG.Response(e.Success, resRoleVO)
}

func EditRole(c *gin.Context)  {
	appG := util.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	if code := RoleIDCheck(id); code != e.Success {
		appG.Response(code, nil)
		return
	}

	roleVO := role_service.RoleVO{}

	code := util.BindAndValid(c, &roleVO) // 数据验证
	if code != e.Success {
		appG.Response(code, nil)
		return
	}

	exists, err := role_service.ExistRoleByRoleName(roleVO.RoleName, id) // 验证用户名是否重复
	if err != nil {
		appG.Response(e.ErrorRoleNameAlreadyExist, nil)
		return
	}
	if exists {
		appG.Response(e.WarningRoleNameAlreadyExist, nil)
		return
	}

	_, err = menu_service.ExistMenuList(roleVO.Menu) //
	if err != nil {
		appG.Response(e.ErrorRoleMenuList, nil)
		return
	}

	resUserVO, err := role_service.EditRole(id, &roleVO)
	if err != nil { // 添加User
		appG.Response(e.ErrorEditUser, nil)
		return
	}

	appG.Response(e.Success, resUserVO)
}
