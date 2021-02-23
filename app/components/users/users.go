package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	userRepo       *UserRepo
	userService    *UserService
	userController *UserController
}

var Service *UserService

func (user *User) NewUser(api fiber.Router, DB *gorm.DB) {

	user.userRepo = NewUserRepo(DB)

	user.userService = NewUserService(user.userRepo)
	Service = NewUserService(user.userRepo)

	user.userController = NewUserController(Service)

	SetUpUsersRoutes(api, user.userController)
}

// SetUpRolesRoutes set routes
func SetUpUsersRoutes(api fiber.Router, userCtrl *UserController) {

	api.Get("/users", userCtrl.Get)
	api.Get("/users/:id", userCtrl.GetByID)
	api.Post("/users", userCtrl.Create)
	api.Put("/users", userCtrl.Update)
	api.Delete("/users/:id", userCtrl.Delete)

}
