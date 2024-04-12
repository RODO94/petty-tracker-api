package controllers

import (
	"github.com/RODO94/petty-tracker-api/initializers"
	"github.com/RODO94/petty-tracker-api/models"
	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context){
	
	var body struct {
		Name string
		UserID string
		HomeID string
		Weighting int16
		Status string
	}

c.Bind(&body)

task := models.Task{Name:body.Name, UserID: body.UserID, HomeID: body.HomeID, Weighting: body.Weighting, Status: "ongoing"}

initializers.DB.Create(&task)

c.JSON(200, gin.H{"message":"Task Created",})
}

func TaskGetOne(c *gin.Context){

	id := c.Param("id")

	var task models.Task
	initializers.DB.Where("deleted_at = ?", nil).First(&task,id)

	c.JSON(200, gin.H{
		"Task":task,
	})
}

func TaskGetAll(c *gin.Context){

	var task models.Task
	initializers.DB.Find(&task)

	c.JSON(200, gin.H{
		"tasks":task,
	})
}

func TaskUpdate(c *gin.Context){
		id := c.Param("id")
		var body struct {
			Name string
			Weighting int16
			Status string
			UserID string
		}
	
		c.Bind(&body)
		var task models.Task
			initializers.DB.First(&task,id )
	
	
	initializers.DB.Model(&task).Updates(models.Task{Name:body.Name, Weighting: body.Weighting, Status: body.Status, UserID: body.UserID})
	
		c.JSON(200, gin.H{
			"task": task,
		})
}

func TaskDelete(c *gin.Context){

	id := c.Param("id")

	initializers.DB.Delete(&models.Task{},id)

	c.JSON(200, gin.H{
		"message":"Task Deleted",
	})
}