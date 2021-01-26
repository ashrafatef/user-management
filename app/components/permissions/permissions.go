package permissions

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Permission struct {
}

func NewPermission(api fiber.Router, DB *gorm.DB) {

	// set repo
	permRepo := NewPermissionRepo(DB)
	fmt.Println(permRepo.db)
	// set service
	premService := NewPermissionService(permRepo)
	// set controller
	permCtrl := NewPermissionController(premService)
	// set Routes
	SetUpPermissionsRoutes(api, permCtrl)
}
