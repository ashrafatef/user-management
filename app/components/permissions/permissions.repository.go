package permissions

import (
	"gorm.io/gorm"
)

type PermissionRepo struct {
	db *gorm.DB
}

// NewUserRepo ..
func NewPermissionRepo(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{
		db: db,
	}
}
