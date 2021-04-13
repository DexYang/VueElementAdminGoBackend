package role_service

import (
	"VueElementAdminGoBackend/models"
	"VueElementAdminGoBackend/pkg/util"
	"VueElementAdminGoBackend/service/menu_service"
)

type RoleVO struct {
	ID    int `json:"id"`
	State int `json:"state"`

	RoleName string `json:"role_name"`
	Remark   string `json:"remark"`

	Menu []int `json:"menu"`
}

func ExistRoleList(ids []uint) (bool, error) {
	return models.ExistRoles(ids)
}

func GetRoleListByIDList(ids []uint) ([]models.Role, error) {
	return models.GetRoleListByIDList(ids)
}

func GetRoleList(offset int, limit int, key string) ([]RoleVO, error) {
	var roles []models.Role

	roles, err := models.GetRoles(offset, limit, key)
	if err != nil {
		return nil, err
	}

	var rolesVO []RoleVO

	if err = util.Mapping(&roles, &rolesVO); err != nil {
		return nil, err
	}

	return rolesVO, nil
}

func GetRoleTotal(key string) (int64, error) {
	count, err := models.GetRoleTotal(key)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ExistRoleByID(id int) (bool, error) {
	return models.ExistRoleByID(id)
}

func GetRole(id int) (*RoleVO, error) {
	role, err := models.GetRole(id)
	if err != nil {
		return nil, err
	}

	menu := role.Menu // 暂存models.Role
	role.Menu = []models.Menu{}

	var roleVO RoleVO

	err = util.Mapping(&role, &roleVO)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(menu); i++ {
		roleVO.Menu = append(roleVO.Menu, menu[i].ID)
	}

	return &roleVO, nil
}

func DeleteRole(id int) (*RoleVO, error) {
	user, err := GetRole(id)
	if err != nil {
		return nil, err
	}

	if err = models.DeleteRole(id); err != nil {
		return nil, err
	}
	return user, nil
}

func ExistRoleByRoleName(roleName string, id int) (bool, error) {
	return models.ExistRoleByRoleName(roleName, id)
}

func AddRole(roleVO *RoleVO) (*RoleVO, error) {
	menuList, err := menu_service.GetMenuListByIDList(roleVO.Menu)
	if err != nil {
		return nil, err
	}

	var role models.Role

	roleVO.Menu = nil
	if util.Mapping(&roleVO, &role) != nil {
		return nil, err
	}
	role.Menu = menuList

	if err := models.AddRole(&role); err != nil {
		return nil, err
	}

	resUserVO, err := GetRole(int(role.ID))
	if err != nil {
		return nil, err
	}

	return resUserVO, nil
}

func EditRole(id int, roleVO *RoleVO) (*RoleVO, error) {
	menuList, err := menu_service.GetMenuListByIDList(roleVO.Menu)
	if err != nil {
		return nil, err
	}

	var role models.Role

	roleVO.Menu = nil
	if util.Mapping(&roleVO, &role) != nil {
		return nil, err
	}
	role.Menu = menuList

	if err := models.UpdateRole(id, &role); err != nil {
		return nil, err
	}

	resUserVO, err := GetRole(id)
	if err != nil {
		return nil, err
	}

	return resUserVO, nil
}
