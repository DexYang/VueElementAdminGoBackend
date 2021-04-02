package menu_service

import (
	"VueElementAdminGoBackend/models"
	"VueElementAdminGoBackend/pkg/util"
)

type MenuVO struct {
	ID uint `json:"id"`

	Name      string `json:"name"`
	Type      int    `json:"type"`
	Title     string `json:"title"`
	Component string `json:"component"`
	Hidden    bool   `json:"hidden"`

	Perms string `json:"perms"`
	Path  string `json:"path"`
	Icon  string `json:"icon"`

	Children []MenuVO `json:"children"`
}

func ExistMenuList(ids []uint) (bool, error) {
	return models.ExistMenuList(ids)
}

func GetMenuListByIDList(ids []uint) ([]models.Menu, error) {
	return models.GetMenuListByIDList(ids)
}

func GetMenu() ([]MenuVO, error) {
	menuList, _ := DFSGetMenu(0)
	var menuVOList []MenuVO

	if err := util.Mapping(&menuList, &menuVOList); err != nil {
		return nil, err
	}
	return menuVOList, nil
}

func DFSGetMenu(parentID int) ([]models.Menu, error) {
	menuList, err := models.GetMenu(parentID)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(menuList); i++ {
		if menuList[i].Type < 2 { // 如果是目录
			menuList[i].Children, _ = DFSGetMenu(int(menuList[i].ID))
		} else {
			menuList[i].Children = []models.Menu{} // no null
		}
	}
	return menuList, nil
}

func SaveMenu(menuVOList []MenuVO) error {
	var menuList []models.Menu

	if err := util.Mapping(&menuVOList, &menuList); err != nil {
		return err
	}

	// 将所有menu的state设为0
	if err := models.SetAllMenuState(0); err != nil {
		return err
	}

	// 递归保存Menu(保存时会将state更新为1)
	if err := DFSSaveMenu(menuList, 0, ""); err != nil {
		return err
	}

	// 将所有state仍为0的menu删除
	if err := models.DeleteAllMenuStateEqZero(); err != nil {
		return err
	}

	return nil
}

func DFSSaveMenu(menuList []models.Menu, parentID uint, perms string) error {
	for i := 0; i < len(menuList); i++ {
		menuList[i].ParentID = parentID
		menuList[i].Order = i
		if perms != "" { // 如果传来的permissionTag不为空，即当前menu是一个按键，需要继承上级页面的permissionTag
			menuList[i].Perms = perms
		}
		// 查询是否存在该ID的menu
		exist, err := models.ExistMenuByID(menuList[i].ID)
		if err != nil {
			return err
		}
		if exist {
			// 更新已有Menu
			if err := models.SaveMenu(&menuList[i]); err != nil {
				return err
			}
		} else {
			// 创建Menu
			if err = models.CreateMenu(&menuList[i]); err != nil {
				return err
			}
		}
		// 如果是目录，则往下递归
		if menuList[i].Type == 0 { // 如果是目录，则permissionTag传为空，即下级菜单的permissionTag不受影响
			if err = DFSSaveMenu(menuList[i].Children, menuList[i].ID, ""); err != nil {
				return err
			}
		} else if menuList[i].Type == 1 { // 如果是页面，则传送当前页面的permissionTag，即该页面下级按键继承当前permissionTag
			if err = DFSSaveMenu(menuList[i].Children, menuList[i].ID, menuList[i].Perms); err != nil {
				return err
			}
		}
	}
	return nil
}

func GetMenuOfUser(username string) ([]MenuVO, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	var roleIds []uint
	for i := 0; i < len(user.Roles); i++ {
		roleIds = append(roleIds, user.Roles[i].ID)
	}

	menuList, _ := DFSGetMenuOfUser(0, roleIds)
	var menuVOList []MenuVO

	if err := util.Mapping(&menuList, &menuVOList); err != nil {
		return nil, err
	}
	return menuVOList, nil
}

func DFSGetMenuOfUser(parentID int, roleIds []uint) ([]models.Menu, error) {
	menuList, err := models.GetMenuByRole(parentID, roleIds)
	if err != nil {
		return nil, err
	}
	if menuList == nil {
		return []models.Menu{}, nil
	}

	for i := 0; i < len(menuList); i++ {
		if menuList[i].Type < 2 { // 如果是目录
			menuList[i].Children, _ = DFSGetMenuOfUser(int(menuList[i].ID), roleIds)
		} else {
			menuList[i].Children = []models.Menu{} // no null
		}
	}
	return menuList, nil
}

func CheckPermission(username string, permissionTag string, permissionTypeList []int) (bool, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	var roleIds []uint
	for i := 0; i < len(user.Roles); i++ {
		roleIds = append(roleIds, user.Roles[i].ID)
	}

	flag, err := models.CheckPermissionByRole(roleIds, permissionTag, permissionTypeList)
	if err != nil {
		return false, err
	}

	return flag, nil
}
