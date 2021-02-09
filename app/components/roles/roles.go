package roles

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Role struct {
}

func NewRole(api fiber.Router, DB *gorm.DB) {

	roleRepo := NewRoleRepo(DB)

	roleService := NewRoleService(roleRepo)

	roleCtrl := NewRoleController(roleService)

	SetUpRolesRoutes(api, roleCtrl)
}

// SetUpRolesRoutes set routes
func SetUpRolesRoutes(api fiber.Router, roleCtrl *RoleContoller) {

	api.Get("/roles", roleCtrl.Get)
	api.Get("/roles/:id", roleCtrl.GetByID)
	api.Post("/roles", roleCtrl.Create)
	api.Put("/roles", roleCtrl.Update)
	api.Delete("/roles/:id", roleCtrl.Delete)

}
