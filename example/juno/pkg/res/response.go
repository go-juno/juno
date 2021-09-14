package res

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-juno/juno/example/juno/pkg/util"
)

type Status string

const (
	Success      Status = "success"
	Failure      Status = "failure"
	ParamCheck   Status = "param check"
	AuthOverdue  Status = "auth overdue"
	NoPermission Status = "no permission"
	SystemError  Status = "system error"
)

type Response struct {
	Status Status      `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func SuccessRes(c *gin.Context, data interface{}) {
	res := Response{
		Status: Success,
		Msg:    nil,
		Data:   data,
	}
	c.JSON(200, res)
}
func FailureRes(c *gin.Context, err error) {
	log.Printf("failure response! error:%+v", err)
	res := Response{
		Status: Failure,
		Msg:    util.Unwrap(err).Error(),
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func ParamCheckRes(c *gin.Context, err error) {
	log.Printf("param check! error:%+v", err)
	res := Response{
		Status: ParamCheck,
		Msg:    util.Unwrap(err).Error(),
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func AuthOverdueRes(c *gin.Context) {
	res := Response{
		Status: AuthOverdue,
		Msg:    "未登录",
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func NoPermissionRes(c *gin.Context) {
	res := Response{
		Status: NoPermission,
		Msg:    "无权限",
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func SystemErrorRes(c *gin.Context, err error) {
	res := Response{
		Status: SystemError,
		Msg:    util.Unwrap(err).Error(),
		Data:   nil,
	}
	c.AbortWithStatusJSON(200, res)
}

func CommonRes(c *gin.Context, status Status, msg string, data interface{}) {
	res := Response{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	c.JSON(200, res)
}
