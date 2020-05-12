package route

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	. "web-02/controller"
	_ "web-02/docs"
)

func InitRoute() *gin.Engine {
	rotuer := gin.Default()
	rotuer.Use(Cors())
	rotuer.GET("/", Index)
	rotuer.GET("/list", ListUser)
	rotuer.GET("/get/:id", Get)
	rotuer.POST("/update", Update)
	rotuer.POST("/add", Add)
	rotuer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return rotuer
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
