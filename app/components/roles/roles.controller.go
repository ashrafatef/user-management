package roles

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RoleContoller struct {
	roleService *RoleService
}

type RoleCreateDTO struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	Permissions    []int  `json:"permissions" validate:"required"`
	OrganizationID int    `json:"organization_id" validate:"required"`
}

type RoleUpdateDTO struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	NewAssign   []int  `json:"newAssign" validate:"required"`
	UnAssign    []int  `json:"UnAssign" validate:"required"`
}

type ErrorResponse struct {
	Failed string
	Tag    string
	Value  string
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

func (roleCtrl *RoleContoller) Update(c *fiber.Ctx) error {

	role := new(RoleUpdateDTO)

	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*role)

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
	errors := ValidateStruct(*role)

	if errors != nil {
		return c.JSON(errors)
	}
	roleCtrl.roleService.Add(*role)
	c.JSON(role)
	return c.SendString("Hello, from create roles!")
}

func ValidateStruct(user interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Failed = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
