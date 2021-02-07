package permissions

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Permission struct {
}

func NewPermission(api fiber.Router, DB *gorm.DB) {
	// set repo
	permRepo := NewPermissionRepo(DB)
	// set service
	premService := NewPermissionService(permRepo)
	// set controller
	permCtrl := NewPermissionController(premService)
	// set Routes
	SetUpPermissionsRoutes(api, permCtrl)
}

func SetUpPermissionsRoutes(api fiber.Router, permCtrl *PermissionController) {

	api.Get("/permissions", permCtrl.Get)

	api.Post("/permissions", permCtrl.Create)

	api.Put("/permissions", permCtrl.Update)

	api.Delete("/permissions/:id", permCtrl.Delete)

}
