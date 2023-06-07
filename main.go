package main

import (
	"github.com/crud_new/controller"
	"github.com/crud_new/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message :" :"POng",
		})
	})
	r.POST("/post", controller.PostCreate)
	r.PUT("/post/:id", controller.PostsUpdate)

	r.GET("/getposts", controller.PostIndex)
	r.GET("/getSinglePost/:id", controller.PostsShow)
	r.DELETE("/delete/:id", controller.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}