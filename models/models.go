package models

import (
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"VueElementAdminGoBackend/pkg/setting"
)

var db *gorm.DB
var tablePrefix string
var sqliteName string

// 扩展gorm.Model 定义
type Model struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	State int
}

func init() {
	var (
		err error
		//dbType, dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database") // 定位到app.ini的database区
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//dbType = sec.Key("TYPE").String()              // 数据库类型mysql
	//dbName = sec.Key("NAME").String()              // 数据库名blog
	//user = sec.Key("USER").String()                // 用户名
	//password = sec.Key("PASSWORD").String()        // 密码
	//host = sec.Key("HOST").String()                // 数据库ip: port
	tablePrefix = sec.Key("TABLE_PREFIX").String() // 表前缀
	sqliteName = sec.Key("SQLITE_NAME").String()   // 表前缀

	db, err = gorm.Open(sqlite.Open(sqliteName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,                       // 表名前缀，`User`表为`t_users`
			SingularTable: true,                              // 使用单数表名，启用该选项后，`User` 表将是`user`
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})
}
