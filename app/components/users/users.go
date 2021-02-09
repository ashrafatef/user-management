package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
}

func NewUser(api fiber.Router, DB *gorm.DB) {

	roleRepo := NewUserRepo(DB)

	roleService := NewUserService(roleRepo)

	roleCtrl := NewUserController(roleService)

	SetUpUsersRoutes(api, roleCtrl)
}

// SetUpRolesRoutes set routes
func SetUpUsersRoutes(api fiber.Router, userCtrl *UserController) {

	api.Get("/users", userCtrl.Get)
	api.Get("/users/:id", userCtrl.GetByID)
	api.Post("/users", userCtrl.Create)
	api.Put("/users", userCtrl.Update)
	api.Delete("/users/:id", userCtrl.Delete)

}
