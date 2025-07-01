package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	r.Group("/api")
	{
	}
	return r
}
