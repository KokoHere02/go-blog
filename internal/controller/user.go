package handler

import (
	"fmt"

	"github.com/KokoHere02/go-blog/config"
	"github.com/KokoHere02/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc     *usecase.UserUsecase
	config *config.Config
}

func NewUserController(uc *usecase.UserUsecase, config *config.Config) *UserController {
	return &UserController{
		uc:     uc,
		config: config,
	}
}

func (uc *UserController) HandlerUser(c *gin.Context) {
	id := c.Param("id")

	fmt.Printf("id: %v\n", id)
}
