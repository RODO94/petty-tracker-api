package main

import (
	"os"

	"github.com/RODO94/petty-tracker-api/controllers"
	"github.com/RODO94/petty-tracker-api/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	server := os.Getenv("SERVER_URL")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user", controllers.UserCreate)
	r.POST("/login", controllers.UserLogin)
	r.GET("/user",controllers.UserGetAll)
	r.GET("/user/:id",controllers.UserGet)
	r.DELETE("/user/:id",controllers.UserDelete)

	r.GET("/home", controllers.HomeGetAll)
	r.GET("/home/:id", controllers.HomeGetOne)
	r.POST("/home", controllers.HomeCreate)
	r.PUT("/home", controllers.HomeUpdate)
	r.DELETE("/home", controllers.HomeDelete)

	r.GET("/task", controllers.TaskGetAll)
	r.GET("/task/:id", controllers.TaskGetOne)
	r.POST("/task", controllers.TaskCreate)
	r.PUT("/task", controllers.TaskUpdate)
	r.DELETE("/task", controllers.TaskDelete)


		r.Run(server) // listen and serve on 0.0.0.0:3000
}
