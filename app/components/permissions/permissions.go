package permissions

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Permission struct {
}

func NewPermission(api fiber.Router, DB *gorm.DB) {
	SetUpPermissionsRoutes(api)
	NewPermissionRepo(DB)
}
