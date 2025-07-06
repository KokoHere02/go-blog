package usecase

import (
	"time"

	"github.com/KokoHere02/go-blog/internal/domain"
)

type UserUsecase struct {
	userRepo       *domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepo *domain.UserRepository, timeout time.Duration) *UserUsecase {
	return &UserUsecase{
		userRepo:       userRepo,
		contextTimeout: timeout,
	}
}

func (u *UserUsecase) GetUserById(id string) (*domain.User, error) {
	return u.userRepo.GetUserById(id)
}
