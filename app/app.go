package app

import (
	"log"
	"userManagementApi/app/components/permissions"
	"userManagementApi/app/components/roles"
	"userManagementApi/app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

//setup function
func SetUp() {
	//load ENV
	loadEnv()
	app := fiber.New()
	api := app.Group("/api")
	// connect to DB
	DB := database.ConnectToDB()
	// load components
	loadComponents(api, DB)
	log.Fatal(app.Listen(":3000"))
}

func loadComponents(api fiber.Router, DB *gorm.DB) {
	permissions.NewPermission(api, DB)
	roles.NewRole(api, DB)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, ", err)
	}
}
