package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// http result
type HttpResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// define erro code
func Error(code int, msg string) HttpResult {
	return HttpResult{code, msg, nil}
}

// http reuslt
func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, HttpResult{code, msg, data})
}

// http ok
func Ok(c *gin.Context) {
	OkWithData(c, nil)
}

// http ok with data
func OkWithData(c *gin.Context, data interface{}) {
	Result(c, OK.Code, OK.Msg, data)
}

// http ok with msg
func OkWithMsg(c *gin.Context, msg string) {
	Result(c, OK.Code, msg, nil)
}

// http fail
func Fail(c *gin.Context) {
	Result(c, ResultError.Code, ResultError.Msg, nil)
}

// http fail with msg
func FailWithMsg(c *gin.Context, code int, msg string) {
	Result(c, code, msg, nil)
}
