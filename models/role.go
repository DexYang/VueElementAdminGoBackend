package models

import (
	"gorm.io/gorm"
)

type Role struct {
	Model

	RoleName string `gorm:"type:varchar(100);not null;index" json:"role_name"`
	Remark   string

	Menu []Menu `gorm:"many2many:role_menu;"` // 用户与角色多对多
}

func ExistRoles(ids []uint) (bool, error) {
	var count int64

	err := db.Model(&User{}).Where("id in (?)", ids).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func GetRoleListByIDList(ids []uint) ([]Role, error) {
	var roles []Role

	err := db.Where("id in (?)", ids).Find(&roles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return roles, nil

}

func GetRoles(offset int, limit int, key string) ([]Role, error) {
	var (
		roles []Role
		err   error
	)

	err = db.
		Where("role_name LIKE ?", "%"+key+"%").
		Or("remark LIKE ? ", "%"+key+"%").
		Offset(offset).
		Limit(limit).
		Find(&roles).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return roles, nil
}

func GetRoleTotal(key string) (int64, error) {
	var count int64

	err := db.
		Model(&Role{}).
		Where("role_name LIKE ?", "%"+key+"%").
		Or("remark LIKE ? ", "%"+key+"%").
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetRole(id int) (*Role, error) {
	var role Role

	err := db.Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&role).Association("Menu").Find(&role.Menu)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &role, nil
}

func ExistRoleByID(id int) (bool, error) {
	var role Role

	err := db.Select("id").Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if role.ID > 0 {
		return true, nil
	}

	return false, nil
}

func DeleteRole(id int) error {
	var role Role

	err := db.Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	_ = db.Model(&role).Association("Menu").Clear()

	if err := db.Unscoped().Where("id = ?", id).Delete(&role).Error; err != nil {
		return err
	}

	return nil
}

func ExistRoleByRoleName(roleName string, id int) (bool, error) {
	var role Role

	err := db.
		Select("id").
		Where(&Role{RoleName: roleName}).
		Not(&Role{Model: Model{ID: id}}).
		First(&role).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if role.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddRole(role *Role) error {
	if err := db.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

func UpdateRole(id int, role *Role) error {
	var originRole Role

	_ = db.Where("id = ?", id).First(&originRole).Association("Menu").Clear()

	if err := db.
		Model(&originRole).
		Where("id = ?", id).
		Updates(&role).
		Error; err != nil {
		return err
	}

	return nil
}
