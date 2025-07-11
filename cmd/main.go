package main

import (
	"time"

	"github.com/KokoHere02/go-blog/config"
	"github.com/KokoHere02/go-blog/internal/db"
	"github.com/KokoHere02/go-blog/internal/router"
	"github.com/gin-gonic/gin"
)

// @title           blog API
// @version         1.0
// @description     This is a blog server celler server.

// @contact.name   API Support
// @contact.url    http://www.go-blog.io/support
// @contact.email  support@blog.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://blog.io/resources/open-api/

func main() {
	route := gin.Default()
	conf := config.NewConfig()
	db := db.NewDB(conf)
	router.Setup(conf, time.Second*10, db, route)
	route.Run(":8080")

	// db, err := gorm.Open(postgres.Open("host=localhost user=admin password=123456 dbname=blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"))
	// if err != nil {
	// 	panic(err)
	// }

	// g := gen.NewGenerator(gen.Config{
	// 	OutPath:      "./dao",
	// 	Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	// 	WithUnitTest: false,
	// })

	// g.UseDB(db)

	// g.GenerateAllTable() // 或 g.GenerateModel("users", "posts")

	// g.Execute()
}
