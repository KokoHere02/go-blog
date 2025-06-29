package service

import (
	"github.com/KokoHere02/go-blog/internal/db"
	"github.com/KokoHere02/go-blog/internal/model"
)

func GetUserById(id string) model.User {
	var user model.User
	db.DB.First(&user, id)
	return user
}
