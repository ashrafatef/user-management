package permissions

import (
	"net/http"
	"strconv"
	"userManagementApi/app/responses"
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

func (permissionCtrl *PermissionController) Get(c *fiber.Ctx) error {
	permissions, Err := permissionCtrl.permService.GetAllPermissions()
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return responses.Success(c, http.StatusOK, permissions)
}

func (permissionCtrl *PermissionController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	Err := permissionCtrl.permService.DeletePermission(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from delete permissions!")
}

func (permissionCtrl *PermissionController) Create(c *fiber.Ctx) error {
	permission := new(PermissionsCreateDTO)
	if err := c.BodyParser(permission); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}

	if errors := validation.ValidateStruct(*permission); errors.Errors != nil {
		return responses.SendError(c, errors)
	}

	Err := permissionCtrl.permService.CreatePermission(permission)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.Status(http.StatusCreated).Send([]byte("Permission Created Successfully"))
}

func (permissionCtrl *PermissionController) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	permission := new(PermissionsUpdateDTO)
	if err := c.BodyParser(permission); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}
	if errors := validation.ValidateStruct(*permission); errors.Errors != nil {
		return responses.SendError(c, errors)
	}
	Err := permissionCtrl.permService.UpdatePermission(permission, id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Updated!")
}
