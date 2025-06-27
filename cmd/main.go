package main

import (
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
	router := gin.Default()
	router.Run(":8080")
}
