package e

var MsgFlags = map[int]string {
	Success :                   "ok",
	Error :                     "fail",
	WarningInvalidParams:       "请求参数错误",

	ErrorCheckAuthTokenFail:    "Token鉴权错误",
	ErrorAuthTokenGenerate:     "Token生成失败",
	ErrorCheckPermission:		"用户权限验证时发生错误",
	ErrorGetUserInfo:			"获取用户权限信息时发生错误",

	WarningWrongAuth:           "用户名或密码错误",
	WarningNotLogin:            "用户未登陆",
	WarningAuthTokenTimeout: 	"Token已过期",
	WarningAuthAlreadyChange:	"账户信息已更改，请重新登录",
	WarningNoPermission:		"没有权限",

	ErrorGetUserList:			"获取用户列表时发生错误",
	ErrorGetUserTotal:			"获取用户数量时发生错误",
	ErrorGetUser:				"获取用户时发生错误",
	ErrorAddUser:				"添加用户时发生错误",
	ErrorUserRoleList:			"验证用户角色列表时发生错误",
	ErrorUserNameAlreadyExist:  "验证用户名是否已存在时发生错误",
	ErrorUserEmailAlreadyExist: "验证邮箱是否已存在时发生错误",
	ErrorDeleteUser:			"删除用户时发生错误",
	ErrorCheckUserExist:		"检查用户是否存在时发生错误",
	ErrorEditUser:				"编辑用户时发生错误",
	WarningUserNameAlreadyExist:"用户名已存在",
	WarningEmailAlreadyExist:	"邮箱已存在",
	WarningUserNotExist:		"用户不存在",

	ErrorGetRoleList:  			"获取角色列表时发生错误",
	ErrorGetRoleTotal:			"获取角色数量时发生错误",
	ErrorCheckRoleExist:		"检查角色是否存在时发生错误",
	ErrorGetRole:				"获取角色时发生错误",
	ErrorDeleteRole:			"删除角色时发生错误",
	ErrorRoleNameAlreadyExist:	"验证角色名是否存在时发生错误",
	ErrorAddRole:				"添加角色时发生错误",
	WarningRoleNotExist:		"角色不存在",
	WarningRoleNameAlreadyExist:"角色名已存在",

	ErrorGetMenu:				"获取目录菜单时发生错误",
	ErrorSaveMenu:				"保存目录菜单时发生错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}

func GetData(code int, data interface{}) interface{} {
	if code == 200 && data != nil {
		return data
	}
	return make(map[string]interface{})
}