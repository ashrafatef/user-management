package roles

import (
	"fmt"
	"net/http"
	"strconv"
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
	var err error
	var id int
	id, err = strconv.Atoi(c.Query("org"))
	roles, err := roleCtrl.roleService.Get(id)
	fmt.Println(roles)
	if err != nil {
		return err
	}
	return c.JSON(roles)
}

func (roleCtrl *RoleContoller) GetByID(c *fiber.Ctx) error {
	var err error
	var id int
	id, err = strconv.Atoi(c.Params("id"))
	role, err := roleCtrl.roleService.GetRoleByID(id)
	if err != nil {
		return err
	}
	return c.JSON(role)
}

func (roleCtrl *RoleContoller) Delete(c *fiber.Ctx) error {
	var err error
	var id int
	id, err = strconv.Atoi(c.Params("id"))
	err = roleCtrl.roleService.DeleteRole(id)
	if err != nil {
		return err
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

	errors := validation.ValidateStruct(*role)

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}
	_, err = roleCtrl.roleService.Update(*role)
	if err != nil {
		return err
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
	errors := validation.ValidateStruct(*role)

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}
	_, err := roleCtrl.roleService.Add(*role)
	if err != nil {
		return err
	}
	return c.SendString("Hello, from create roles!")
}
