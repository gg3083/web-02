package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"web-02/model"
)

func init_data() ([]model.User) {
	user := make([]model.User, 2)
	user[0] = model.User{1, "张珊", "2019-01-02", true}
	user[1] = model.User{2, "张珊2", "2019-11-02", false}
	return user
}

func ListUser(c *gin.Context) {
	users := init_data()
	c.JSON(200,gin.H{"data": users})
}


func Get(c *gin.Context) {
	id, _ := strconv.Atoi( c.Param("id"))

	users := init_data()
	user := model.User{}
	for k, v := range users {
		if id == v.Id {
			fmt.Println("遍历:",k)
			user = v
			continue
		}
	}
	c.JSON(200,gin.H{"data":user})
}

func Update(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("ctx.Request.body: %v", string(data))

	params := fmt.Sprintf("%v", string(data))
	fmt.Println(params)
	fmt.Println("========================")
	keyBytes, _ := json.Marshal(data)
	fmt.Println(string(keyBytes))
	fmt.Println("========================")
	user := model.User{}
	par := []byte(params)
	json.Unmarshal(par,&user)
	fmt.Println("========================")
	fmt.Println("%+v\n",user)
	fmt.Println("%#v\n",user)
	//fmt.Println(keyBytes)
	c.JSON(200,gin.H{"data":user})
}