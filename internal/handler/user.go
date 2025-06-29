package handler

import (
	"github.com/KokoHere02/go-blog/internal/service"
	"github.com/gin-gonic/gin"
)

func HandlerUser(c *gin.Context) {
	id := c.Param("id")
	user := service.GetUserById(id)
	c.JSON(200, user)
}
