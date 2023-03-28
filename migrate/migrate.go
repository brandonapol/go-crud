package main

import (
	"github.com/brandonapol/go-crud/initializers"
	"github.com/brandonapol/go-crud/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
