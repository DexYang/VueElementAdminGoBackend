package user_service

import (
	"VueElementAdminGoBackend/models"
	"VueElementAdminGoBackend/pkg/util"
	"VueElementAdminGoBackend/service/role_service"
)

type UserVO struct {
	ID    uint `json:"id"`
	State int  `json:"state"`

	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`

	Roles []int `json:"roles"` // 用户与角色多对多
}

func GetUserList(offset int, limit int, key string) ([]UserVO, error) {
	var users []models.User

	users, err := models.GetUsers(offset, limit, key)
	if err != nil {
		return nil, err
	}

	var usersVO []UserVO

	if err = util.Mapping(&users, &usersVO); err != nil {
		return nil, err
	}

	return usersVO, nil
}

func GetUserTotal(key string) (int64, error) {
	count, err := models.GetUserTotal(key)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ExistUserByID(id int) (bool, error) {
	return models.ExistUserByID(id)
}

func GetUser(id int) (*UserVO, error) {
	user, err := models.GetUser(id)
	if err != nil {
		return nil, err
	}

	roleModels := user.Roles // 暂存models.Role
	user.Roles = []models.Role{}

	var userVO UserVO

	err = util.Mapping(&user, &userVO)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(roleModels); i++ {
		userVO.Roles = append(userVO.Roles, roleModels[i].ID)
	}

	return &userVO, nil
}

func ExistUserByUsername(username string, id int) (bool, error) {
	return models.ExistUserByUsername(username, id)
}

func ExistUserByEmail(email string, id int) (bool, error) {
	return models.ExistUserByEmail(email, id)
}

func DeleteUser(id int) (*UserVO, error) {

	user, err := models.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	var userVO UserVO

	err = util.Mapping(&user, &userVO)
	if err != nil {
		return nil, err
	}

	return &userVO, nil
}

type UserRequest struct {
	ID    uint `json:"id"`
	State int  `json:"state"`

	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`

	Roles []uint `json:"roles"` // 用户与角色多对多
}

func EditUser(id int, userRq *UserRequest) (*UserVO, error) {
	roleList, err := role_service.GetRoleListByIDList(userRq.Roles)
	if err != nil {
		return nil, err
	}

	var user models.User

	userRq.Roles = nil
	if util.Mapping(&userRq, &user) != nil {
		return nil, err
	}
	user.Roles = roleList

	if err := models.UpdateUser(id, &user); err != nil {
		return nil, err
	}

	resUserVO, err := GetUser(id)
	if err != nil {
		return nil, err
	}

	return resUserVO, nil
}

func AddUser(userRq *UserRequest) (*UserVO, error) {
	roleList, err := role_service.GetRoleListByIDList(userRq.Roles)
	if err != nil {
		return nil, err
	}

	var user models.User

	userRq.Roles = nil
	if util.Mapping(&userRq, &user) != nil {
		return nil, err
	}
	user.Roles = roleList

	if err := models.AddUser(&user); err != nil {
		return nil, err
	}

	resUserVO, err := GetUser(int(user.ID))
	if err != nil {
		return nil, err
	}

	return resUserVO, nil
}
