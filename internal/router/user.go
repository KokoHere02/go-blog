package router

import (
	"time"

	"github.com/KokoHere02/go-blog/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouter(config *config.Config, timeout time.Duration, db *gorm.DB, gin *gin.RouterGroup) {

}
