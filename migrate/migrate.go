package main

import (
	"gin_api/core/models"
	"gin_api/initializer"
)

func init() {
	initializer.LoadEnvVaraibles()
	initializer.ConnectionToDB()
}

func main() {
	// Drop existing tables (optional, remove if you want to keep existing data)
	initializer.DB.Migrator().DropTable(&models.Post{})
	initializer.DB.Migrator().DropTable(&models.User{})

	// Create event store table
	initializer.DB.AutoMigrate(&models.EventStore{})
}
