package controllers

import (
	"github.com/RODO94/petty-tracker-api/initializers"
	"github.com/RODO94/petty-tracker-api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(c *gin.Context){
	var body struct {
		Name string
		Email string
		Password string
	}

	c.Bind(&body)

	u1 := uuid.NewV4()
	
	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)

	user := models.User{ID:u1,Name:body.Name, Email:body.Email, Password: string(hashed)}


	if err != nil {
		c.Status(400)
		return
	}

	initializers.DB.Create(&user)


	c.JSON(200, gin.H{
		"message":"User Created",
	})
}

func UserLogin(c *gin.Context){

	var body struct {
		Email string
		Password string
	}

	c.Bind(&body)

	var user models.User
	initializers.DB.Where("Email = ?",body.Email).First(&user)


	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if result == nil {
		c.JSON(200, gin.H{
			"message":"Login Succesful",
		})
	} else {
		c.JSON(400, gin.H{
			"message":"Login Unsuccesful",
		})
	}

}

func UserGet(c *gin.Context){

	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user,id)

	c.JSON(200, gin.H{
		"user":user,
	})
}

func UserGetAll(c *gin.Context){

	var user models.User
	initializers.DB.Find(&user)

	c.JSON(200, gin.H{
		"user":user,
	})
}