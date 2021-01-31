package roles

import (
	"fmt"
	"net/http"
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
	fmt.Println("from get")
	roles := roleCtrl.roleService.Get()
	fmt.Println(roles)
	return c.SendString("Hello, from get roles!")
}

// Update update
func (roleCtrl *RoleContoller) Update(c *fiber.Ctx) error {

	role := new(RoleUpdateDTO)

	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validation.ValidateStruct(*role)

	if errors != nil {
		return c.JSON(errors)
	}
	roleCtrl.roleService.Update(*role)
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
		// return c.JSON(errors)
		return fiber.NewError(http.StatusMethodNotAllowed)

	}
	r, err := roleCtrl.roleService.Add(*role)
	if err != nil {
		return err
	}
	c.JSON(r)
	return c.SendString("Hello, from create roles!")
}
