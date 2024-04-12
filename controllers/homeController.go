package controllers

import (
	"github.com/RODO94/petty-tracker-api/initializers"
	"github.com/RODO94/petty-tracker-api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func HomeCreate(c *gin.Context){
	var body struct {
		Name string
		UserID string
	}

	c.Bind(&body)

	u1 := uuid.NewV4()

	home := models.Home{ID:u1, Name:body.Name, UserID:body.UserID}

	initializers.DB.Create(&home)

	c.JSON(200, gin.H{
		"message":"Home created",
		"home":home,
	})
}

func HomeGetOne(c *gin.Context){

	id := c.Param("id")

	var home models.Home
	initializers.DB.Where("deleted_at = ?", nil).First(&home,id)

	c.JSON(200, gin.H{
		"Home":home,
	})
}

func HomeGetAll(c *gin.Context){

	var home []models.Home
	initializers.DB.Find(&home)

	c.JSON(200, gin.H{
		"homes":home,
	})
}

func HomeUpdate(c *gin.Context){
		id := c.Param("id")
		var body struct {
			Name string
		}
	
		c.Bind(&body)
		var home models.Home
			initializers.DB.First(&home,id )
	
	
	initializers.DB.Model(&home).Updates(models.Home{Name:body.Name})
	
		c.JSON(200, gin.H{
			"home": home,
		})
}

func HomeDelete(c *gin.Context){

	id := c.Param("id")

	initializers.DB.Delete(&models.Home{},id)

	c.JSON(200, gin.H{
		"message":"Home Deleted",
	})
}