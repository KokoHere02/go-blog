package db

import (
	"fmt"

	"github.com/KokoHere02/go-blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	dsn     = "host=localhost user=admin password=123456 dbname=blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// )

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Name,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	fmt.Println("✅ GORM 已连接数据库")
}
