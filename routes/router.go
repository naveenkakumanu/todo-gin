package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naveenkakumanu/todoapp/controllers"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Create User
	router.POST("/user/create", controllers.CreateUser)
	router.GET("/user/read/:id", controllers.ReadUser)
	router.PUT("/user/update/:id", controllers.UpdateUser)
	router.DELETE("/user/delete/:id", controllers.DeleteUser)

	// Create Task
	router.POST("/app/create", controllers.Create)
	router.GET("/app/read/:id", controllers.Read)
	router.PUT("/app/update/:id", controllers.Update)
	router.DELETE("/app/delete/:id", controllers.Delete)
	return router
}
