package app

import (
	"fmt"
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
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			fmt.Println("FROOOOOOOOOOOOOOOOOOOOOOOOOM EROOOOOOOOOOOOOOOOORr")
			fmt.Println("FROOOOOOOOOOOOOOOOOOOOOOOOOM EROOOOOOOOOOOOOOOOORr", ctx)
			fmt.Println("FROOOOOOOOOOOOOOOOOOOOOOOOOM EROOOOOOOOOOOOOOOOORr", err)
			code := fiber.StatusInternalServerError

			// Retreive the custom statuscode if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				fmt.Println(e.Message)
				fmt.Println(e.Code)
				code = e.Code
			}

			errorFormate := fiber.Map{
				"meta": fiber.Map{
					"Version": "1.0",
				},
				"data": fiber.Map{
					"code":  code,
					"error": err.Error(),
				},
			}
			// Send custom error page
			// err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			// if err != nil {
			// In case the SendFile fails
			return ctx.Status(code).JSON(errorFormate)
			// }

			// Return from handler
			return nil
		},
	})
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
