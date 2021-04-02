package models

import (
	"VueElementAdminGoBackend/pkg/setting"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	State int `gorm:"default:0"`

	Username string `gorm:"type:varchar(100);not null;index"`
	Password string
	Email    string `gorm:"type:varchar(100);not null;index"`
	Mobile   string
	IsAdmin  bool

	Roles []Role `gorm:"many2many:user_role;"` // 用户与角色多对多
}

func HmacSha256(key, data string) string {
	hash := hmac.New(sha256.New, []byte(key)) //创建对应的sha256哈希加密算法
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func Encrypt(data string) string {
	sec, err := setting.Cfg.GetSection("salt") // 定位到app.ini的database区
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	salt := sec.Key("SALT").String() // 数据库类型mysql
	return HmacSha256(salt, data)
}

/**
  验证用户名密码
*/
func CheckUser(username, password string) bool {
	var user User
	db.Select("id").Where(User{Username: username, Password: Encrypt(password)}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func GetUsers(offset int, limit int, key string) ([]User, error) {
	var (
		users []User
		err   error
	)

	err = db.
		Where("username LIKE ?", "%"+key+"%").
		Or("email LIKE ? ", "%"+key+"%").
		Offset(offset).
		Limit(limit).
		Find(&users).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

func GetUserTotal(key string) (int64, error) {
	var count int64

	err := db.
		Model(&User{}).
		Where("username LIKE ?", "%"+key+"%").
		Or("email LIKE ? ", "%"+key+"%").
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func ExistUserByID(id int) (bool, error) {
	var user User

	err := db.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetUser(id int) (*User, error) {
	var user User

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&user).Association("Roles").Find(&user.Roles)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func ExistUserByUsername(username string, id int) (bool, error) {
	var user User

	err := db.
		Select("id").
		Where(&User{Username: username}).
		Not(&User{ID: uint(id)}).
		First(&user).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistUserByEmail(email string, id int) (bool, error) {
	var user User

	err := db.
		Select("id").
		Where(User{Email: email}).
		Not(&User{ID: uint(id)}).
		First(&user).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func DeleteUser(id int) (*User, error) {
	var user User

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user.IsAdmin == true {
		return nil, nil
	}

	_ = db.Model(&user).Association("Roles").Clear()

	if err := db.Unscoped().Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func AddUser(user *User) error {
	user.Password = Encrypt(user.Password)
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(id int, user *User) error {
	var originUser User

	_ = db.Where("id = ?", id).First(&originUser).Association("Roles").Clear()

	originUser.Username = user.Username
	if user.Password != "" {
		originUser.Password = Encrypt(user.Password)
	}
	originUser.Email = user.Email
	originUser.Mobile = user.Mobile
	originUser.State = user.State
	originUser.Roles = user.Roles

	if err := db.Save(&originUser).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User

	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&user).Association("Roles").Find(&user.Roles)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
