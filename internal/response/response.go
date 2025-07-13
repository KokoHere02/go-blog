package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Code    int    `json:"code"`           // 状态码
	Message string `json:"message"`        // 提示信息
	Data    T      `json:"data,omitempty"` // 返回的数据
}

func Ok[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Response[T]{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response[any]{
		Code:    code,
		Message: msg,
	})
}
