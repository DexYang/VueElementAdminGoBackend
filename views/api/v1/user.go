package v1

import (
	"github.com/DeluxeYang/GinProject/pkg/e"
	"github.com/DeluxeYang/GinProject/pkg/util"
	"github.com/DeluxeYang/GinProject/service/menu_service"
	"github.com/DeluxeYang/GinProject/service/role_service"
	"github.com/DeluxeYang/GinProject/service/user_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)


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
		"list": users,
		"total": total,
		"page": page,
		"limit": limit,
	})
}

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

func AddUser(c *gin.Context) {
	appG := util.Gin{C: c}

	userVO := user_service.UserVO{}

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
	}

	appG.Response(e.Success, user)
}

func UserIDCheck(id int) int{
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

func EditUser(c *gin.Context)  {
	appG := util.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	if code := UserIDCheck(id); code != e.Success {
		appG.Response(code, nil)
		return
	}

	userVO := user_service.UserVO{}

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
