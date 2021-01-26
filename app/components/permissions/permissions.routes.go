package permissions

import (
	"github.com/gofiber/fiber/v2"
)

// SetUpPermissionsRoutes setup all permission routes
func SetUpPermissionsRoutes(api fiber.Router, permCtrl *PermissionController) {

	api.Get("/permissions", permCtrl.Get)

	api.Post("/permissions", permCtrl.Create)

	api.Put("/permissions", permCtrl.Update)

}
