package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
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
	// router := gin.Default()
	// router.Run(":8080")

	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=123456 dbname=blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dao",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest: false,
	})

	g.UseDB(db)

	g.GenerateAllTable() // 或 g.GenerateModel("users", "posts")

	g.Execute()
}
