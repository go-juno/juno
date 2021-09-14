package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"juno/pkg/res"
)

//处理错误的中间件
func ErrMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			log.Printf("err:%+v\n stack:%s", err, string(debug.Stack()))
			res := res.Response{
				Status: res.SystemError,
				Msg:    fmt.Sprintf("%v", r),
				Data:   nil,
			}
			c.AbortWithStatusJSON(http.StatusOK, res)
		}
	}()
	c.Next()

}
