package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/go-juno/juno/example/juno/pkg/res"
)

//处理错误的中间件
func ErrMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("err:%+v\n stack:%s", r, string(debug.Stack()))
			response := res.Response{
				Status: res.SystemError,
				Msg:    fmt.Sprintf("%v", r),
				Data:   nil,
			}
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}
	}()
	c.Next()
}
