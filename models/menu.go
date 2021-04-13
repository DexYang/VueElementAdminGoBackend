package models

import (
	"errors"
	"gorm.io/gorm"
)

type Menu struct {
	Model

	ParentID int `gorm:"default:0" json:"parent_id"`

	Name      string `json:"name"`
	Type      int    `json:"type"`
	Title     string `json:"title"`
	Component string `json:"component"`
	Hidden    bool   `json:"hidden"`

	Perms string `json:"perms"`

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

	result := db.Where("parent_id = ?", parentID).Order("'order'").Find(&menu)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, err
	}

	return menu, nil
}

func SetAllMenuState(state int) error {
	return db.Model(&Menu{}).Where("1=1").Update("state", state).Error
}

func ExistMenuByID(id int) (bool, error) {
	var menu Menu
	result := db.Where("id = ?", id).First(&menu)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}

	if result.RowsAffected <= 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, nil
}

func SaveMenu(menu *Menu) error {
	menu.State = 1
	if err := db.Model(&menu).Omit("CreateAt").Updates(&menu).Error; err != nil {
		return err
	}
	return nil
}

func CreateMenu(menu *Menu) error {
	menu.State = 1
	if err := db.Omit("ID", "CreateAt").Create(&menu).Error; err != nil {
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
		_ = db.Model(&menus[i]).Association("Role").Clear()
	}

	if err := db.Unscoped().Where("state = 0").Delete(&Menu{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistMenuList(ids []int) (bool, error) {
	var count int64

	err := db.Model(&Menu{}).Where("id in (?)", ids).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func GetMenuListByIDList(ids []int) ([]Menu, error) {
	var menuList []Menu

	err := db.Where("id in (?)", ids).Find(&menuList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return menuList, nil
}

func CheckPermissionByRole(roleIds []int, permissionTag string, permissionTypeList []int) (bool, error) {
	type result struct {
		ID int
	}

	var res result

	err := db.Raw("SELECT menu.id FROM "+
		tablePrefix+"menu as menu, "+
		tablePrefix+"role_menu as role_menu "+
		"WHERE role_menu.role_id in (?) "+
		"AND menu.id = role_menu.menu_id "+
		"AND menu.perms = ? "+
		"AND menu.type in (?)", roleIds, permissionTag, permissionTypeList).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if res != (result{}) {
		return true, nil
	}

	return false, nil
}

func GetMenuByRole(parentID int, roleIds []int) ([]Menu, error) {
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
