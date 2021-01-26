package permissions

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PermissionDTO struct {
	Name string `json:"name"`
}

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
	prems := permCont.permService.GetAllPermissions()
	fmt.Println(prems)
	return c.SendString("Hello, from get permissions!")
}

func (permCont *PermissionController) Create(c *fiber.Ctx) error {
	permission := new(Permissions)
	if err := c.BodyParser(permission); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}
	fmt.Println(permission)
	permCont.permService.Create(permission)
	return c.Status(http.StatusCreated).Send([]byte("Permission Created Successfully"))
}

func (permCont *PermissionController) Update(c *fiber.Ctx) error {
	permission := new(Permissions)
	if err := c.BodyParser(permission); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}
	permCont.permService.UpdatePermission(permission)
	return c.SendString("Updated!")
}
