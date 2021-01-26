package permissions

import (
	"fmt"

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

func (permRepo *PermissionRepo) CreatePermissionRepo(permission *Permissions) {
	fmt.Println("in Repo Permission")
	permRepo.db.Create(&permission)
}

func (permRepo *PermissionRepo) GetAll() *[]Permissions {
	var permission []Permissions

	fmt.Println("permission")
	permRepo.db.Find(&permission)
	fmt.Println(permission)
	return &permission
}

func (permRepo *PermissionRepo) UpdatePermissionRepo(permission *Permissions) {
	fmt.Println("in Repo Permission")
	permRepo.db.Model(&permission).Update("name", permission.Name)
}
