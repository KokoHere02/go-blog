package router

import (
	"time"

	"github.com/KokoHere02/go-blog/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(config *config.Config, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicGroup := gin.Group("/api/user")
	NewUserRouter(config, timeout, db, publicGroup)

	gin.Group("/api/article")
}
