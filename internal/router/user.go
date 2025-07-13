package router

import (
	"time"

	"github.com/KokoHere02/go-blog/config"
	handler "github.com/KokoHere02/go-blog/internal/controller"
	"github.com/KokoHere02/go-blog/internal/domain"
	"github.com/KokoHere02/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouter(config *config.Config, timeout time.Duration, db *gorm.DB, gin *gin.RouterGroup) {
	ur := domain.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur, time.Second*5)
	uc := handler.NewUserController(uu, config)
	gin.GET("/:id", uc.QueryUser)
}
