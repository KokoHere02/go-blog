package usecase

import (
	"github.com/KokoHere02/go-blog/internal/domain"
)

type ArticleService interface {
	GetArticleById(id string) (*domain.Article, error)
}
