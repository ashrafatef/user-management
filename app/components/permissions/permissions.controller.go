package permissions

import (
	"fmt"
	"net/http"
	"strconv"
	"userManagementApi/app/validation"

	"github.com/gofiber/fiber/v2"
)

type PermissionController struct {
	permService *PermissionService
}

func NewPermissionController(permService *PermissionService) *PermissionController {
	return &PermissionController{
		permService: permService,
	}
}

func (permCont *PermissionController) Get(c *fiber.Ctx) error {
	fmt.Println("from get")
	perms, err := permCont.permService.GetAllPermissions()
	fmt.Println(perms)
	if err != nil {
		return err
	}
	return c.JSON(perms)
}

func (permCont *PermissionController) Delete(c *fiber.Ctx) error {
	var err error
	var id int
	id, err = strconv.Atoi(c.Params("id"))
	err = permCont.permService.DeletePermission(id)
	if err != nil {
		return err
	}
	return c.SendString("Hello, from delete permissions!")
}

func (permCont *PermissionController) Create(c *fiber.Ctx) error {
	permission := new(PermissionsCreateDTO)
	if err := c.BodyParser(permission); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}

	errors := validation.ValidateStruct(*permission)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}
	fmt.Println(permission)

	err := permCont.permService.CreatePermission(permission)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).Send([]byte("Permission Created Successfully"))
}

func (permCont *PermissionController) Update(c *fiber.Ctx) error {
	permission := new(PermissionsUpdateDTO)
	if err := c.BodyParser(permission); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}
	errors := validation.ValidateStruct(*permission)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}
	err := permCont.permService.UpdatePermission(permission)
	if err != nil {
		return err
	}
	return c.SendString("Updated!")
}
