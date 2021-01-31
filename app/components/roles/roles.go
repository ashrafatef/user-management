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

	SetUpPermissionsRoutes(api, roleCtrl)
}

func SetUpPermissionsRoutes(api fiber.Router, roleCtrl *RoleContoller) {

	api.Get("/roles", roleCtrl.Get)

	api.Post("/roles", roleCtrl.Create)

	api.Put("/roles", roleCtrl.Update)

}
