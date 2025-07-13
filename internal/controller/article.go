package controller

import (
	"net/http"

	"github.com/KokoHere02/go-blog/config"
	"github.com/KokoHere02/go-blog/internal/response"
	"github.com/KokoHere02/go-blog/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleUsecase *usecase.ArticleUsecase
	config         *config.Config
}

func NewArticleController(articleUsecase *usecase.ArticleUsecase, config *config.Config) *ArticleController {
	return &ArticleController{
		articleUsecase: articleUsecase,
		config:         config,
	}
}

func (c *ArticleController) QueryArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")
	article, err := c.articleUsecase.GetArticleById(articleID)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Ok(ctx, article)
}
