package main

import (
	"github.com/asifDurran/Golang-crud/initializers"
	"github.com/asifDurran/Golang-crud/models"
)

func init (){
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main(){
	 // Migrate the schema
	 initializers.DB.AutoMigrate(&models.Post{})
	 initializers.DB.AutoMigrate(&models.User{})
}