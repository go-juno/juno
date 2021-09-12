package handle

import (
	"github.com/gin-gonic/gin"
	"juno/api/http/schema"
	"juno/api/http/serialize"
	"juno/internal/endpoint"
	"juno/pkg/res"
	"golang.org/x/xerrors"
)

func GreetingBluePrint(v1 *gin.RouterGroup, endpoints *endpoint.Endpoints) {
	r := v1.Group("/greeting-router")
	r.GET("list", getGreetingList(endpoints))
	r.GET("all", getGreetingAll(endpoints))
	r.GET("", getGreetingDetail(endpoints))
	r.POST("", createGreeting(endpoints))
	r.PUT("", updateGreeting(endpoints))
	r.DELETE("", deleteGreeting(endpoints))
}

/**
@api {GET} /api/greeting-router/list 获取greeting列表
@apiVersion 1.0.0
@apiName getGreetingList
@apiGroup greeting
@apiParam {number} page_index 页码
@apiParam {number} page_size 每页大小
@apiSuccess {string} status 状态码
@apiSuccess {string} msg 返回的信息
@apiSuccess {list} data 返回的数据
@apiSuccess {number} data.id id
@apiSuccess {string} data.created_at 创建时间
@apiSuccess {string} data.updated_at 更新时间
@apiSuccessExample Success-Response:
{"status":"success","msg":null,"data":{"items":[{"id":2,"created_at":"2021-07-16 03:53:48","updated_at":"2021-07-16 03:53:48"}],"total":1}}
@apiErrorExample Error-Response:
{"status": "failure", "data": null, "msg": "失败"}
**/
func getGreetingList(endpoints *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		var scheme schema.GetGreetingList
		err := c.ShouldBind(&scheme)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.ParamCheckRes(c, err)
			return
		}
		req := scheme.Transform()
		result, err := endpoints.GetGreetingListEndpoint(c, req)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.FailureRes(c, err)
			return
		}
		res.SuccessRes(c, serialize.GetGreetingListTransform(result))
	}
}

/**
@api {GET} /api/greeting-router/all 获取全部greeting数据
@apiVersion 1.0.0
@apiName getGreetingAll
@apiGroup greeting
@apiSuccess {string} status 状态码
@apiSuccess {string} msg 返回的信息
@apiSuccess {list} data 返回的数据
@apiSuccess {number} data.id id
@apiSuccess {string} data.created_at 创建时间
@apiSuccess {string} data.updated_at 更新时间
@apiSuccessExample Success-Response:
{"status":"success","msg":null,"data":[{"id":2,"created_at":"2021-07-16 03:53:48","updated_at":"2021-07-16 03:53:48"}]}
@apiErrorExample Error-Response:
{"status": "failure", "data": null, "message": "失败"}
**/
func getGreetingAll(endpoints *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		var scheme schema.GetGreetingAll
		err := c.ShouldBind(&scheme)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.ParamCheckRes(c, err)
			return
		}
		req := scheme.Transform()
		result, err := endpoints.GetGreetingAllEndpoint(c, req)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.FailureRes(c, err)
			return
		}
		res.SuccessRes(c, serialize.GetGreetingAllTransform(result))
	}
}

/**
@api {GET} /api/greeting-router/detail 获取greeting详情
@apiVersion 1.0.0
@apiName getGreetingDetail
@apiGroup greeting
@apiSuccess {string} status 状态码
@apiSuccess {string} msg 返回的信息
@apiSuccess {list} data 返回的数据
@apiSuccess {number} data.id id
@apiSuccess {string} data.created_at 创建时间
@apiSuccess {string} data.updated_at 更新时间
@apiSuccessExample Success-Response:
{"status":"success","msg":null,"data":{"items":[{"id":2,"created_at":"2021-07-16 03:53:48","updated_at":"2021-07-16 03:53:48"}],"total":1}}
@apiErrorExample Error-Response:
{"status": "failure", "data": null, "message": "失败"}
**/
func getGreetingDetail(endpoints *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		var scheme schema.GetGreetingDetail
		err := c.ShouldBind(&scheme)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.ParamCheckRes(c, err)
			return
		}
		req := scheme.Transform()
		result, err := endpoints.GetGreetingDetailEndpoint(c, req)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.FailureRes(c, err)
			return
		}
		res.SuccessRes(c, serialize.GetGreetingDetailTransform(result))
	}
}

/**
@api {POST} /api/greeting-router 创建greeting
@apiVersion 1.0.0
@apiName createGreeting
@apiGroup greeting
@apiSuccess {string} status 状态码
@apiSuccess {string} msg 返回的信息
@apiSuccess {dict} data 返回的数据
@apiSuccessExample Success-Response:
{"status":"success","msg":null,"data":{}}
@apiErrorExample Error-Response:
{"status": "failure", "data": null, "message": "失败"}
**/
func createGreeting(endpoints *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		var scheme schema.CreateGreeting
		err := c.ShouldBind(&scheme)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.ParamCheckRes(c, err)
			return
		}
		req := scheme.Transform()
		result, err := endpoints.CreateGreetingEndpoint(c, req)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.FailureRes(c, err)
			return
		}
		res.SuccessRes(c, serialize.CreateGreetingTransform(result))
	}
}

/**
@api {PUT} /api/greeting-router 更新greeting
@apiVersion 1.0.0
@apiName updateGreeting
@apiGroup greeting
@apiParam {number} id id
@apiSuccess {string} status 状态码
@apiSuccess {string} msg 返回的信息
@apiSuccess {dict} data 返回的数据
@apiSuccessExample Success-Response:
{"status":"success","msg":null,"data":{}}
@apiErrorExample Error-Response:
{"status": "failure", "data": null, "message": "失败"}
**/
func updateGreeting(endpoints *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		var scheme schema.UpdateGreeting
		err := c.ShouldBind(&scheme)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.ParamCheckRes(c, err)
			return
		}
		req := scheme.Transform()
		result, err := endpoints.UpdateGreetingEndpoint(c, req)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.FailureRes(c, err)
			return
		}
		res.SuccessRes(c, serialize.UpdateGreetingTransform(result))
	}
}

/**
@api {DELETE} /api/greeting-router 删除greeting
@apiVersion 1.0.0
@apiName deleteGreeting
@apiGroup greeting
@apiParam {number} id id
@apiSuccess {string} status 状态码
@apiSuccess {string} msg 返回的信息
@apiSuccess {dict} data 返回的数据
@apiSuccessExample Success-Response:
{"status":"success","msg":null,"data":{}}
@apiErrorExample Error-Response:
{"status": "failure", "data": null, "message": "失败"}
**/
func deleteGreeting(endpoints *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		var scheme schema.DeleteGreeting
		err := c.ShouldBind(&scheme)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.ParamCheckRes(c, err)
			return
		}
		req := scheme.Transform()
		result, err := endpoints.DeleteGreetingEndpoint(c, req)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			res.FailureRes(c, err)
			return
		}
		res.SuccessRes(c, serialize.DeleteGreetingTransform(result))
	}
}
