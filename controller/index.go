package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	a := fmt.Sprint("sss")
	fmt.Println(a)
	c.JSON(http.StatusOK,gin.H{"message":"main"})
}
