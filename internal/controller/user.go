package controller

import (
	"fmt"

	"github.com/KokoHere02/go-blog/config"
	"github.com/KokoHere02/go-blog/internal/response"
	"github.com/KokoHere02/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase *usecase.UserUsecase
	config      *config.Config
}

func NewUserController(uc *usecase.UserUsecase, config *config.Config) *UserController {
	return &UserController{
		userUsecase: uc,
		config:      config,
	}
}

func (uc *UserController) QueryUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("id: %v\n", id)
	user, err := uc.userUsecase.GetUserById(id)
	if err != nil {
		response.Fail(c, 1, "internal server error")
		return
	}
	response.Ok(c, user)
}
