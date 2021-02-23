package roles

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Role struct {
	roleRepo      *RoleRepo
	roleService   *RoleService
	roleContoller *RoleContoller
}

func (role *Role) NewRole(api fiber.Router, DB *gorm.DB) {

	role.roleRepo = NewRoleRepo(DB)

	role.roleService = NewRoleService(role.roleRepo)

	role.roleContoller = NewRoleController(role.roleService)

	SetUpRolesRoutes(api, role.roleContoller)
}

// SetUpRolesRoutes set routes
func SetUpRolesRoutes(api fiber.Router, roleCtrl *RoleContoller) {

	api.Get("/roles", roleCtrl.Get)
	api.Get("/roles/:id", roleCtrl.GetByID)
	api.Post("/roles", roleCtrl.Create)
	api.Put("/roles", roleCtrl.Update)
	api.Delete("/roles/:id", roleCtrl.Delete)

}
