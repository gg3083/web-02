package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags 首页
// @Success 200 {object} model.User
// @Failure 400 {object} model.User
// @Router / [get]
func Index(c *gin.Context) {
	formate := "2006-01-02 15:04:05"
	now := time.Now()
	fmt.Println(now)
	local, _ := time.LoadLocation("Local") //输入参数"UTC"，等同于""

	main := fmt.Sprintf("现在是 %s", now.In(local).Format(formate))
	c.JSON(http.StatusOK, gin.H{"message": main})
}
