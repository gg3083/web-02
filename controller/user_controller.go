package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web-02/model"
	"web-02/model/e"
)

func initData() []model.User {
	user := make([]model.User, 2)
	user[0] = model.User{1, "张珊", "2019-01-02", true}
	user[1] = model.User{2, "张珊2", "2019-11-02", false}
	return user
}

// @Summary 获取用户信息
// @Description 获取所有用户信息
// @Tags 用户信息
// @accept json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /list [get]
func ListUser(c *gin.Context) {
	var (
		appG = model.Gin{C: c}
		//form   getChannelInfoListForm
		errMsg string
	)
	users, err := model.User{}.UserList()
	if err != nil {
		errMsg = err.Error()
		appG.Response(e.ERROR, e.INVALID_PARAMS, errMsg, users)
		return
	}
	appG.Response(e.SUCCESS, e.SUCCESS, errMsg, users)

}

// @Summary 获取单个用户信息
// @Description 根据Id获取所有用户信息
// @Tags 用户信息
// @accept json
// @Produce  json
// @param id path int true "ID"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /get/{id} [get]
func Get(c *gin.Context) {
	var (
		appG   = model.Gin{C: c}
		errMsg string
	)
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := model.User{}.GetUser(id)

	if err != nil {
		errMsg = err.Error()
		appG.Response(e.SUCCESS, e.ERROR, errMsg, user)
		return
	}
	appG.Response(e.SUCCESS, e.SUCCESS, errMsg, user)
}

type saveUserFrom struct {
	Id       int    `json:"id"			`
	UserName string `json:"userName"     valid:"Required"`
	Birthday string `json:"birthday"     valid:"Required"`
	Sex      bool   `json:"sex"          valid:"Required"`
}

// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags 用户信息
// @accept json
// @Produce  json
// @param request body model.User true "id代理 userName承兑人账户名 birthday承兑人姓名 sex电话号码"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /update [post]
func Update(c *gin.Context) {
	var (
		appG   = model.Gin{C: c}
		form   saveUserFrom
		errMsg string
	)
	httpCode, errCode := model.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg, nil)
		return
	}

	user := model.User{
		Id:       form.Id,
		UserName: form.UserName,
		Sex:      form.Sex,
		Birthday: form.Birthday,
	}
	res, err := model.User{}.SaveUser(user)

	if err != nil {
		errMsg = err.Error()
		appG.Response(httpCode, errCode, errMsg, res)
		return
	}
	appG.Response(httpCode, errCode, errMsg, res)
}

// @Summary 添加用户信息
// @Description 添加所有用户信息
// @Tags 用户信息
// @accept json
// @Produce  json
// @param request body model.User true "id代理 userName承兑人账户名 birthday承兑人姓名 sex电话号码"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /add [post]
func Add(c *gin.Context) {
	var (
		appG   = model.Gin{C: c}
		form   saveUserFrom
		errMsg string
	)
	httpCode, errCode := model.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errMsg, nil)
		return
	}

	user := model.User{
		Id:       form.Id,
		UserName: form.UserName,
		Sex:      form.Sex,
		Birthday: form.Birthday,
	}
	res, err := model.User{}.SaveUser(user)

	if err != nil {
		errMsg = err.Error()
		appG.Response(httpCode, errCode, errMsg, res)
		return
	}
	appG.Response(httpCode, errCode, errMsg, res)
}
