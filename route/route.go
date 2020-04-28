package route

import (
	"github.com/gin-gonic/gin"
	. "web-02/controller"
)

func InitRoute()  *gin.Engine{
	rotuer := gin.Default()
	rotuer.GET("/",Index)
	rotuer.GET("/list",ListUser)
	rotuer.GET("/get/:id",Get)
	rotuer.POST("/update",Update)
	return rotuer
}
