package main

import (
	"github.com/RODO94/petty-tracker-api/initializers"
	"github.com/gin-gonic/gin"
)

func init(){

initializers.ConnectToDB()
initializers.LoadEnvVariables()
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
		r.Run() // listen and serve on 0.0.0.0:3000
}