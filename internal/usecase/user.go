package usecase

import (
	"github.com/KokoHere02/go-blog/internal/domain"
)

type UserService interface {
	GetUserById(id string) (*domain.User, error)
}
