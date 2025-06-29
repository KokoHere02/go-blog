package router

import (
	"github.com/KokoHere02/go-blog/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserGroupRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("/:id", handler.HandlerUser)
	}
}
