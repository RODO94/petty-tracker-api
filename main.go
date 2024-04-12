package main

import (
	"github.com/RODO94/petty-tracker-api/controllers"
	"github.com/RODO94/petty-tracker-api/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/user", controllers.UserCreate)
	r.POST("/login", controllers.UserLogin)

	r.GET("/home", controllers.HomeGetAll)
	r.GET("/home/:id", controllers.HomeGetOne)
	r.POST("/home", controllers.HomeCreate)
	r.PUT("/home", controllers.HomeUpdate)
	r.DELETE("/home", controllers.HomeDelete)

	
		r.Run() // listen and serve on 0.0.0.0:3000
}
