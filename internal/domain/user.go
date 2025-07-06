package domain

import (
	"time"

	"github.com/KokoHere02/go-blog/internal/db"
	"gorm.io/gorm"
)

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Nickname   string    `json:"nickname"`
	Avatar_url string    `json:"avatar_url"`
	Bio        string    `json:"bio"`
	Is_admin   bool      `json:"is_admin"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.NewDB(),
	}
}

func (user *UserRepository) GetUserById(id string) (*User, error) {
	userDb := user.db
	var userDo User
	err := userDb.First(&userDo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &userDo, nil
}
