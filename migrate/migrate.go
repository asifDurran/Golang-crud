package main

import (
	"github.com/crud_new/initializers"
	"github.com/crud_new/models"
)

func init (){
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main(){
	 // Migrate the schema
	 initializers.DB.AutoMigrate(&models.Post{})
}