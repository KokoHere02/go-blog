package db

import (
	"fmt"

	"github.com/KokoHere02/go-blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// var (
// 	dsn     = "host=localhost user=admin password=123456 dbname=blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// )


func NewDB() *gorm.DB {
	cfg := config.New() 
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		"disable",
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // ✅ 不自动加 s
		},
	})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	fmt.Println("✅ GORM 已连接数据库")
	return DB
}
