package v1

import (
	"VueElementAdminGoBackend/pkg/e"
	"VueElementAdminGoBackend/pkg/util"
	"VueElementAdminGoBackend/service/menu_service"
	"VueElementAdminGoBackend/service/role_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Summary Get role list
// @Produce  json
// @Tags roles
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param key query string true "key"
// @Success 200 {object} util.Response "{"code":200,"data":{"list":[], "limit":10, "page":1, "total":100},"message":"ok"}"
// @Router /api/v1/roles [get]
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
		"list":  roles,
		"total": total,
		"page":  page,
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

// @Summary Get a single role
// @Produce  json
// @Tags roles
// @Param id path int true "ID"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "role_name":"x", "remark": "", "menu": []},"message":"ok"}"
// @Router /api/v1/roles/{id} [get]
func GetRole(c *gin.Context) {
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

// @Summary Delete a single role
// @Produce  json
// @Tags roles
// @Param id path int true "ID"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "role_name":"x", "remark": "", "menu": []},"message":"ok"}"
// @Router /api/v1/roles/{id} [delete]
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

// @Summary Create a single role
// @Produce  json
// @Tags roles
// @Param role_name body string true "role_name"
// @Param remark body string true "remark"
// @Param menu body array true "menu"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "role_name":"x", "remark": "", "menu": []},"message":"ok"}"
// @Router /api/v1/roles [post]
func AddRole(c *gin.Context) {
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

// @Summary Update a single role
// @Produce  json
// @Tags roles
// @Param id path int true "ID"
// @Param role_name body string true "role_name"
// @Param remark body string true "remark"
// @Param menu body array true "menu"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "role_name":"x", "remark": "", "menu": []},"message":"ok"}"
// @Router /api/v1/roles/{id} [put]
func EditRole(c *gin.Context) {
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
