package roles

import (
	"net/http"
	"strconv"
	"userManagementApi/app/responses"
	"userManagementApi/app/validation"

	"github.com/gofiber/fiber/v2"
)

type RoleContoller struct {
	roleService *RoleService
}

func NewRoleController(roleService *RoleService) *RoleContoller {
	return &RoleContoller{
		roleService: roleService,
	}
}

func (roleCtrl *RoleContoller) Get(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Query("org"))
	roles, Err := roleCtrl.roleService.Get(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return responses.Success(c, http.StatusOK, roles)
}

func (roleCtrl *RoleContoller) GetByID(c *fiber.Ctx) error {
	// var err error
	// var id int
	id, _ := strconv.Atoi(c.Params("id"))
	role, Err := roleCtrl.roleService.GetRoleByID(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return responses.Success(c, http.StatusOK, role)
}

func (roleCtrl *RoleContoller) Delete(c *fiber.Ctx) error {
	// var err error
	// var id int
	id, _ := strconv.Atoi(c.Params("id"))
	Err := roleCtrl.roleService.DeleteRole(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from get roles!")
}

// Update update
func (roleCtrl *RoleContoller) Update(c *fiber.Ctx) error {
	var err error
	role := new(RoleUpdateDTO)

	if err = c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if errors := validation.ValidateStruct(*role); errors.Errors != nil {
		return responses.SendError(c, errors)
	}
	_, Err := roleCtrl.roleService.Update(*role)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from update roles!")

}

func (roleCtrl *RoleContoller) Create(c *fiber.Ctx) error {
	role := new(RoleCreateDTO)

	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := validation.ValidateStruct(*role); errors.Errors != nil {
		return responses.SendError(c, errors)
	}

	_, Err := roleCtrl.roleService.Add(*role)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from create roles!")
}
