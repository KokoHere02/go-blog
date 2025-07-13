package usecase

import (
	"time"

	"github.com/KokoHere02/go-blog/internal/domain"
)

type ArticleUsecase struct {
	articleRepo    *domain.ArticleRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(articleRepo *domain.ArticleRepository, timeout time.Duration) *ArticleUsecase {
	return &ArticleUsecase{
		articleRepo:    articleRepo,
		contextTimeout: timeout,
	}
}

func (u *ArticleUsecase) GetArticleById(id string) (*domain.Article, error) {
	return u.articleRepo.GetArticleById(id)
}
