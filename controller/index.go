package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags 首页
// @Success 200 {object} model.User
// @Failure 400 {object} model.User
// @Router / [get]
func Index(c *gin.Context) {
	a := fmt.Sprint("main")
	fmt.Println(a)
	c.JSON(http.StatusOK, gin.H{"message": "main"})
}
