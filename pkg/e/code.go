package e

const (
	Success       = 200
	Error         = 500
	WarningInvalidParams   		= 10000

	ErrorAuthTokenGenerate 		= 10001
	ErrorCheckAuthTokenFail    	= 10003
	ErrorCheckPermission		= 10005
	ErrorGetUserInfo			= 10007

	WarningWrongAuth            = 10002
	WarningNotLogin			   	= 10004
	WarningAuthTokenTimeout		= 10006
	WarningAuthAlreadyChange	= 10008
	WarningNoPermission 		= 10010

	ErrorGetUserList 			= 11001
	ErrorGetUserTotal			= 11003
	ErrorGetUser				= 11005
	ErrorAddUser 				= 11007
	ErrorUserRoleList			= 11009
	ErrorUserNameAlreadyExist	= 11011
	ErrorUserEmailAlreadyExist	= 11013
	ErrorDeleteUser				= 11015
	ErrorCheckUserExist			= 11017
	ErrorEditUser				= 11019

	WarningUserNameAlreadyExist	= 11002
	WarningEmailAlreadyExist	= 11004
	WarningUserNotExist			= 11006


	ErrorGetRoleList			= 12001
	ErrorGetRoleTotal			= 12003
	ErrorCheckRoleExist			= 12005
	ErrorGetRole				= 12007
	ErrorDeleteRole				= 12009
	ErrorRoleNameAlreadyExist	= 12011
	ErrorRoleMenuList			= 12013
	ErrorAddRole				= 12015

	WarningRoleNotExist			= 12002
	WarningRoleNameAlreadyExist	= 12004

	ErrorGetMenu				= 13001
	ErrorSaveMenu				= 13003
)