package controller

import (
	"github.com/crud_new/initializers"
	"github.com/crud_new/models"
	"github.com/gin-gonic/gin"
)

func PostCreate (c *gin.Context) {
   // Get data off req body
   var body struct {
	Body string
	Title string
   }
   c.Bind(&body)

   //Create a post 
   post := models.Post{Title: body.Title,Body: body.Body}

    result := initializers.DB.Create(&post) 
	
	if result.Error != nil {
		c.Status(400)
		return
	}

	 c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostIndex(c *gin.Context){
	//Get the posts 
   var posts []models.Post
   initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}


func PostsShow(c *gin.Context){
	// Get ID from params
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostsUpdate(c *gin.Context){
	// Get ID from param
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)
	//Find the post were updating
	var post models.Post
	initializers.DB.Find(&post,id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Responsd
	c.JSON(200, gin.H{
		"Post": post,
	})


}

func PostDelete (c *gin.Context){
	id := c.Param("id")
	var post models.Post
	initializers.DB.Delete(&post, id)


	c.JSON(200, gin.H{
		"Post": "Delete",
	})
}