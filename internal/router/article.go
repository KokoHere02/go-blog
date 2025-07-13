package router

import (
	"time"

	"github.com/KokoHere02/go-blog/config"
	"github.com/KokoHere02/go-blog/internal/controller"
	"github.com/KokoHere02/go-blog/internal/domain"
	"github.com/KokoHere02/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewArticleRouter(config *config.Config, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	articleRepo := domain.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepo, timeout)
	articleController := controller.NewArticleController(articleUsecase, config)
	gin.GET("/article/:id", articleController.QueryArticle)
}
