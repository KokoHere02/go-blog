package handler

import (
	"time"

	"github.com/KokoHere02/go-blog/internal/domain"
	"github.com/KokoHere02/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

func HandlerUser(c *gin.Context) {
	id := c.Param("id")
	ur := domain.NewUserRepository()
	uc := usecase.NewUserUsecase(ur, time.Second*5)
	user, err := uc.GetUserById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
