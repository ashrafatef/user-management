package permissions

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Permission struct {
	permissionRepo       *PermissionRepo
	permissionService    *PermissionService
	permissionController *PermissionController
}

func (permission *Permission) NewPermission(api fiber.Router, DB *gorm.DB) {
	permission.permissionRepo = NewPermissionRepo(DB)
	permission.permissionService = NewPermissionService(permission.permissionRepo)
	permission.permissionController = NewPermissionController(permission.permissionService)
	SetUpPermissionsRoutes(api, permission.permissionController)
}

func SetUpPermissionsRoutes(api fiber.Router, permCtrl *PermissionController) {

	api.Get("/permissions", permCtrl.Get)

	api.Post("/permissions", permCtrl.Create)

	api.Put("/permissions", permCtrl.Update)

	api.Delete("/permissions/:id", permCtrl.Delete)

}
