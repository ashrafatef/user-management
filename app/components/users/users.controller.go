package users

import (
	"net/http"
	"strconv"
	"userManagementApi/app/responses"
	"userManagementApi/app/validation"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userCtrl *UserController) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("org"))
	users, Err := userCtrl.userService.Get(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return responses.Success(c, http.StatusOK, users)
}

func (userCtrl *UserController) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, Err := userCtrl.userService.GetUserByID(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return responses.Success(c, http.StatusOK, user)
}

func (userCtrl *UserController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	Err := userCtrl.userService.Delete(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from get roles!")
}

func (userCtrl *UserController) Update(c *fiber.Ctx) error {
	var err error
	user := new(UserUpdateDTO)

	if err = c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := validation.ValidateStruct(*user); errors.Errors != nil {
		return responses.SendError(c, errors)

	}
	_, Err := userCtrl.userService.Update(*user)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from update roles!")
}

func (userCtrl *UserController) Create(c *fiber.Ctx) error {
	user := new(UserCreateDTO)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if errors := validation.ValidateStruct(*user); errors.Errors != nil {
		return responses.SendError(c, errors)
	}
	_, Err := userCtrl.userService.Add(*user)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from create roles!")
}
