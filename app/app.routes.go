package app

import (
	"userManagementApi/app/components/permissions"

	"github.com/gofiber/fiber/v2"
)

// type AppRoutes struct{
// 	app *fiber.App
// }

// func NewAppRoutes(app *fiber.App)  {
// 	app =  app
// }

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	permissions.SetUpPermissionsRoutes(api)
}
