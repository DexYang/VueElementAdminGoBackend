package models

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	Model

	ParentID uint `gorm:"default:0" json:"parent_id"`

	MenuName  string `json:"menu_name"`
	MenuType  int    `json:"menu_type"`
	Remark    string `json:"remark"`
	Component string `json:"component"`

	PermissionTag string `json:"permission_tag"`

	Path  string `json:"path"`
	Icon  string `json:"icon"`
	Order int    `json:"order"`

	Children []Menu `gorm:"-"`

	Role []Role `gorm:"many2many:role_menu;"` // 用户与角色多对多
}

func GetMenu(parentID int) ([]Menu, error) {
	var (
		menu []Menu
		err  error
	)

	err = db.Where("parent_id = ?", parentID).Order("order").Find(&menu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return menu, nil
}

func SetAllMenuState(state int) error {
	return db.Model(&Menu{}).Update("state", state).Error
}

func ExistMenuByID(id uint) (bool, error) {
	var menu Menu
	err := db.Select("id").Where("id = ?", id).First(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if menu.ID > 0 {
		return true, nil
	}

	return false, nil
}

func SaveMenu(menu *Menu) error {
	menu.State = 1
	if err := db.Save(&menu).Error; err != nil {
		return err
	}
	return nil
}

func CreateMenu(menu *Menu) error {
	menu.State = 1
	if err := db.Create(&menu).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAllMenuStateEqZero() error {
	var menus []Menu
	err := db.Where("state = 0").Find(&menus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	for i := 0; i < len(menus); i++ {
		db.Model(&menus[i]).Association("Role").Clear()
	}

	if err := db.Unscoped().Where("state = 0").Delete(&Menu{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistMenuList(ids []uint) (bool, error) {
	var count int

	err := db.Model(&Menu{}).Where("id in (?)", ids).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func GetMenuListByIDList(ids []uint) ([]Menu, error) {
	var menuList []Menu

	err := db.Where("id in (?)", ids).Find(&menuList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return menuList, nil
}

func CheckPermissionByRole(roleIds []uint, permissionTag string, permissionTypeList []int) (bool, error) {
	type result struct {
		ID uint
	}

	var res result

	err := db.Raw("SELECT menu.id FROM "+
		tablePrefix+"menu as menu, "+
		tablePrefix+"role_menu as role_menu "+
		"WHERE role_menu.role_id in (?) "+
		"AND menu.id = role_menu.menu_id "+
		"AND menu.permission_tag = ? "+
		"AND menu.menu_type in (?)", roleIds, permissionTag, permissionTypeList).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if res != (result{}) {
		return true, nil
	}

	return false, nil
}

func GetMenuByRole(parentID int, roleIds []uint) ([]Menu, error) {
	var menu []Menu

	err := db.Raw("SELECT DISTINCT menu.* FROM "+
		tablePrefix+"menu as menu INNER JOIN "+
		tablePrefix+"role_menu as role_menu ON "+
		"menu.id = role_menu.menu_id WHERE "+
		"role_menu.role_id in (?) AND menu.parent_id = ?", roleIds, parentID).Order("order").Scan(&menu).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return menu, nil
}
