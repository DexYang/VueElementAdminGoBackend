package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/DeluxeYang/GinProject/pkg/setting"
)

var db *gorm.DB
var tablePrefix string

// 扩展gorm.Model 定义
type Model struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	State int
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database") // 定位到app.ini的database区
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String() // 数据库类型mysql
	dbName = sec.Key("NAME").String() // 数据库名blog
	user = sec.Key("USER").String() // 用户名
	password = sec.Key("PASSWORD").String() // 密码
	host = sec.Key("HOST").String() // 数据库ip: port
	tablePrefix = sec.Key("TABLE_PREFIX").String() // 表前缀

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)) // db 为数据库连接，类型为*gorm.DB

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName
	} // 默认表名处理器

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func init() {
	db.AutoMigrate(&User{})
}


func CloseDB() {
	defer db.Close()
}