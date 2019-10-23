package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"VueElementAdminGoBackend/models"
	"VueElementAdminGoBackend/pkg/setting"
)

var db *gorm.DB

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database") // 定位到app.ini的database区
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()              // 数据库类型mysql
	dbName = sec.Key("NAME").String()              // 数据库名blog
	user = sec.Key("USER").String()                // 用户名
	password = sec.Key("PASSWORD").String()        // 密码
	host = sec.Key("HOST").String()                // 数据库ip: port
	tablePrefix = sec.Key("TABLE_PREFIX").String() // 表前缀

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)) // db 为数据库连接，类型为*gorm.DB

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	} // 默认表名处理器

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func main() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Menu{})
}

func CloseDB() {
	defer db.Close()
}
