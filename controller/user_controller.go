package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"web-02/model"
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
// @Success 200 {object} model.User
// @Failure 400 {object} model.User
// @Router /list [get]
func ListUser(c *gin.Context) {
	users, err := model.User{}.UserList()
	c.JSON(200, gin.H{"data": users, "errMsg": err})
}

// @Summary 获取单个用户信息
// @Description 根据Id获取所有用户信息
// @Tags 用户信息
// @accept json
// @Produce  json
// @param id path int true "ID"
// @Success 200 {object} model.User
// @Failure 500 {object} model.User
// @Router /get/{id} [get]
func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := model.User{}.GetUser(id)

	c.JSON(200, gin.H{"data": user, "errMsg": err})
}

// @Summary 获取单个用户信息
// @Description 根据Id获取所有用户信息
// @Tags 用户信息
// @accept json
// @Produce  json
// @param request body model.User true "id代理 userName承兑人账户名 birthday承兑人姓名 sex电话号码"
// @Success 200 {object} model.User
// @Failure 500 {object} model.User
// @Router /update [post]
func Update(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	params := fmt.Sprintf("%v", string(data))
	fmt.Println(params)
	var getBody model.User
	if err := json.Unmarshal(data, &getBody); err != nil {
		fmt.Println(" transcode get, body Unmarshal error:%v", err)
		return
	}
	fmt.Println("========================")
	fmt.Println(getBody)
	res, err := model.User{}.UpdateUser(getBody)
	//fmt.Println(keyBytes)
	c.JSON(200, gin.H{"data": res, "errMsg": err})
}
