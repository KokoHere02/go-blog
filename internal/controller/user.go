package controller

import (
	"fmt"
	"strconv"

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

func parseID(s string) (int64, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string cannot convert to int64")
	}
	return strconv.ParseInt(s, 10, 64)
}

func (uc *UserController) QueryUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("id: %v\n", id)
	intId, err := parseID(id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		response.Fail(c, 500, err.Error())
		return
	}

	user, err := uc.userUsecase.GetUserById(intId)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		response.Fail(c, 500, err.Error())
		return
	}
	response.Ok(c, user)
}
