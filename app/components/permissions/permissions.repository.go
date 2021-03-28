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

func (permRepo *PermissionRepo) CreatePermissionRepo(permission *Permissions) error {
	res := permRepo.db.Create(&permission)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (permRepo *PermissionRepo) GetAllPermissionsRepo() ([]Permissions, error) {
	var permissions []Permissions

	res := permRepo.db.Find(&permissions)
	if res.Error != nil {
		return []Permissions{}, res.Error
	}
	return permissions, nil
}

func (permRepo *PermissionRepo) UpdatePermissionRepo(permission *Permissions) error {
	res := permRepo.db.Model(&permission).Update("name", permission.Name)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (permRepo *PermissionRepo) DeletePermissionRepo(permID int) error {
	res := permRepo.db.Delete(&Permissions{}, permID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (permRepo *PermissionRepo) CheckPermissionAssigning(permID int) (int64, error) {
	var id int
	res := permRepo.db.Raw("SELECT id FROM permissions_roles WHERE permission_id=?", permID).Scan(&id)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
