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

func ListUser(c *gin.Context) {
	users, err := model.User{}.UserList()
	c.JSON(200, gin.H{"data": users, "errMsg": err})
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := model.User{}.GetUser(id)

	c.JSON(200, gin.H{"data": user, "errMsg": err})
}

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
