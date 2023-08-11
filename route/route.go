package route

import (
	"github.com/gin-gonic/gin"

	"github.com/rakshitgondwal/swagger-api-doc/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/", controller.GetUsers)
		user.POST("/", controller.CreateUser)
		user.GET("/:id", controller.GetUserByID)
	}
	return r
}