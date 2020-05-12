package model

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"web-02/model/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, errMsg string, data interface{}) {

	g.C.JSON(httpCode, Response{
		Code:   errCode,
		Msg:    e.GetMsg(errCode),
		ErrMsg: errMsg,
		Data:   data,
	})
	return
}

// ErrorResp 错误返回值
func ErrorResp(g *Gin, code int, errMsg string) {
	resp(g, code, errMsg, nil)
}

// SuccessResp 正确返回值
func SuccessResp(g *Gin, data interface{}) {
	resp(g, e.SUCCESS, "", data)
}

// resp 返回
func resp(g *Gin, code int, errMsg string, data ...interface{}) {
	resp := Response{
		Code:   code,
		Msg:    e.GetMsg(code),
		ErrMsg: errMsg,
		Data:   data,
	}
	g.C.JSON(http.StatusOK, resp)
}

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.ShouldBindJSON(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Println(err.Key, err.Message)
	}

	return
}
