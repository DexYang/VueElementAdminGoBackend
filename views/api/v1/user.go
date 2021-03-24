package v1

import (
	"VueElementAdminGoBackend/pkg/e"
	"VueElementAdminGoBackend/pkg/util"
	"VueElementAdminGoBackend/service/menu_service"
	"VueElementAdminGoBackend/service/role_service"
	"VueElementAdminGoBackend/service/user_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Summary Get user info
// @Produce  json
// @Success 200 {object} util.Response "{"code":200,"data":{"menus":[], "username":"x"},"message":"ok"}"
// @Router /api/v1/info [get]
func GetUserInfo(c *gin.Context) {
	appG := util.Gin{C: c}
	data := make(map[string]interface{})

	usernameValue, exists := c.Get("username")
	if exists {
		username := com.ToStr(usernameValue)

		menu, err := menu_service.GetMenuOfUser(username)
		if err != nil {
			appG.Response(e.ErrorGetUserInfo, nil)
			return
		}
		data["menus"] = menu
		data["username"] = username
	} else {
		appG.Response(e.ErrorGetUserInfo, nil)
		return
	}

	appG.Response(e.Success, data)
}

// @Summary Get user list
// @Produce  json
// @Tags users
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param key query string true "key"
// @Success 200 {object} util.Response "{"code":200,"data":{"list":[], "limit":10, "page":1, "total":100},"message":"ok"}"
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	appG := util.Gin{C: c}

	offset, page, limit := util.GetPage(c)
	key := c.DefaultQuery("key", "")

	users, err := user_service.GetUserList(offset, limit, key)
	if err != nil {
		appG.Response(e.ErrorGetUserList, nil)
		return
	}

	total, err := user_service.GetUserTotal(key)
	if err != nil {
		appG.Response(e.ErrorGetUserTotal, nil)
		return
	}

	appG.Response(e.Success, map[string]interface{}{
		"list":  users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// @Summary Get a single user
// @Produce  json
// @Tags users
// @Param id path int true "ID"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "username":"x", "email": "", "mobile": "", "state":1, "roles": []},"message":"ok"}"
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {
	appG := util.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	if code := UserIDCheck(id); code != e.Success {
		appG.Response(code, nil)
		return
	}

	user, err := user_service.GetUser(id)
	if err != nil {
		appG.Response(e.ErrorGetUser, nil)
		return
	}

	appG.Response(e.Success, user)
}

// @Summary Create a single user
// @Produce  json
// @Tags users
// @Param username body string true "username"
// @Param email body string true "email"
// @Param roles body array true "roles"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "username":"x", "email": "", "mobile": "", "state":1, "roles": []},"message":"ok"}"
// @Router /api/v1/users [post]
func AddUser(c *gin.Context) {
	appG := util.Gin{C: c}

	userVO := user_service.UserRequest{}

	code := util.BindAndValid(c, &userVO) // 数据验证
	if code != e.Success {
		appG.Response(code, nil)
		return
	}

	exists, err := user_service.ExistUserByUsername(userVO.Username, 0) // 验证用户名是否重复
	if err != nil {
		appG.Response(e.ErrorUserNameAlreadyExist, nil)
		return
	}
	if exists {
		appG.Response(e.WarningUserNameAlreadyExist, nil)
		return
	}

	exists, err = user_service.ExistUserByEmail(userVO.Email, 0) // 验证邮箱是否重复
	if err != nil {
		appG.Response(e.ErrorUserEmailAlreadyExist, nil)
		return
	}
	if exists {
		appG.Response(e.WarningEmailAlreadyExist, nil)
		return
	}

	_, err = role_service.ExistRoleList(userVO.Roles) // 验证角色ID列表是否有错
	if err != nil {
		appG.Response(e.ErrorUserRoleList, nil)
		return
	}

	resUserVO, err := user_service.AddUser(&userVO)
	if err != nil { // 添加User
		appG.Response(e.ErrorAddUser, nil)
		return
	}

	appG.Response(e.Success, resUserVO)
}

// @Summary Delete a single user
// @Produce  json
// @Tags users
// @Param id path int true "ID"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "username":"x", "email": "", "mobile": "", "state":1, "roles": []},"message":"ok"}"
// @Router /api/v1/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	appG := util.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	if code := UserIDCheck(id); code != e.Success {
		appG.Response(code, nil)
		return
	}

	user, err := user_service.DeleteUser(id)
	if err != nil {
		appG.Response(e.ErrorDeleteUser, nil)
		return
	} else if user == nil {
		appG.Response(e.WarningCannotDeleteAdmin, nil)
		return
	}

	appG.Response(e.Success, user)
}

func UserIDCheck(id int) int {
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		return e.WarningInvalidParams
	}

	exists, err := user_service.ExistUserByID(id)
	if err != nil {
		return e.ErrorCheckUserExist
	}
	if !exists {
		return e.WarningUserNotExist
	}

	return e.Success
}

// @Summary Update a single user
// @Produce  json
// @Tags users
// @Param id path int true "ID"
// @Param username body string true "username"
// @Param email body string true "email"
// @Param roles body array true "roles"
// @Success 200 {object} util.Response "{"code":200,"data":{"id":1, "username":"x", "email": "", "mobile": "", "state":1, "roles": []},"message":"ok"}"
// @Router /api/v1/users/{id} [put]
func EditUser(c *gin.Context) {
	appG := util.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	if code := UserIDCheck(id); code != e.Success {
		appG.Response(code, nil)
		return
	}

	userVO := user_service.UserRequest{}

	code := util.BindAndValid(c, &userVO) // 数据验证
	if code != e.Success {
		appG.Response(code, nil)
		return
	}

	exists, err := user_service.ExistUserByUsername(userVO.Username, id) // 验证用户名是否重复
	if err != nil {
		appG.Response(e.ErrorUserNameAlreadyExist, nil)
		return
	}
	if exists {
		appG.Response(e.WarningUserNameAlreadyExist, nil)
		return
	}

	exists, err = user_service.ExistUserByEmail(userVO.Email, id) // 验证邮箱是否重复
	if err != nil {
		appG.Response(e.ErrorUserEmailAlreadyExist, nil)
		return
	}
	if exists {
		appG.Response(e.WarningEmailAlreadyExist, nil)
		return
	}

	_, err = role_service.ExistRoleList(userVO.Roles) // 验证角色ID列表是否有错
	if err != nil {
		appG.Response(e.ErrorUserRoleList, nil)
		return
	}

	resUserVO, err := user_service.EditUser(id, &userVO)
	if err != nil { // 添加User
		appG.Response(e.ErrorEditUser, nil)
		return
	}

	appG.Response(e.Success, resUserVO)
}
