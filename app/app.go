package app

import (
	"log"
	"os"
	"userManagementApi/app/components/bots"
	"userManagementApi/app/components/permissions"
	"userManagementApi/app/components/roles"
	"userManagementApi/app/components/users"
	"userManagementApi/app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

//setup function
func SetUp() {
	//load ENV
	loadEnv()
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// Retreive the custom statuscode if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			errorFormate := fiber.Map{
				"meta": fiber.Map{
					"Version": "1.0",
				},
				"errors": fiber.Map{
					"code":  code,
					"error": err.Error(),
				},
			}
			return ctx.Status(code).JSON(errorFormate)
		},
	})

	app.Use(cors.New())
	api := app.Group("/api")
	// connect to DB
	DB := database.ConnectToDB()
	// load components
	loadComponents(api, DB)
	log.Fatal(app.Listen(os.Getenv("PORT")))
}

func loadComponents(api fiber.Router, DB *gorm.DB) {
	permission := new(permissions.Permission)
	roles := new(roles.Role)
	users := new(users.User)
	bots := new(bots.Bot)

	permission.NewPermission(api, DB)
	roles.NewRole(api, DB)
	users.NewUser(api, DB)
	bots.NewBot(api, DB)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, ", err)
	}
}
