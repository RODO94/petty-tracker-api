// package controllers

// import (
// 	"github.com/gin-gonic/gin"
// )

// func PostsCreate(c *gin.Context) {
// 	// Get data off req body
// 	var body struct {
// 		Body string
// 		Title string
// 	}

// 	c.Bind(&body)

// 	// Create Post
// 	post := models.Post{Title: body.Title, Body: body.Body}

// 	result := initializers.DB.Create(&post)


// 	if result.Error != nil {
// 		c.Status(400)
// 		return
// 	}
// 	// Return Post
// 	c.JSON(200, gin.H{
// 		"post": post,
// 	})
// }

// func PostsIndex(c *gin.Context){

// 	// Get the Posts
// 	// Get all records
// 	var posts []models.Post
// 		initializers.DB.Find(&posts)

// 	// Response with them

// 	c.JSON(200, gin.H{
// 		"post": posts,
// 	})
// }

// func PostsShow(c *gin.Context){
// 	// Get the Posts
// 	// Get id off url
// 	id := c.Param("id")
// 	// Get all records
// 	var post models.Post
// 		initializers.DB.First(&post,id )

// 	// Response with them

// 	c.JSON(200, gin.H{
// 		"post": post,
// 	})
// }

// func PostsUpdate(c *gin.Context){
// 	// Get id off url
// 	id := c.Param("id")
// 	// Get the data off req body
// 	var body struct {
// 		Body string
// 		Title string
// 	}

// 	c.Bind(&body)
// 	// find the post from DB
// 	var post models.Post
// 		initializers.DB.First(&post,id )

// // Update post

// initializers.DB.Model(&post).Updates(models.Post{Title:body.Title, Body:body.Body,})

// 	// Response with post
// 	c.JSON(200, gin.H{
// 		"post": post,
// 	})
// }

// func PostsDelete(c *gin.Context){
// 	// Get the ID off the URL
// 	id := c.Param("id")

// 	// Delete the posts
// 	initializers.DB.Delete(&models.Post{},id)

// 	// Response
// 	c.Status(200)

//