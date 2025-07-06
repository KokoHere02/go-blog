package db

import (
	"fmt"
	"log"

	"github.com/KokoHere02/go-blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDB() *gorm.DB {
	cfg := config.NewConfig()
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
