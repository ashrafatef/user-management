package roles

import (
	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

// RoleRepoInterface role repo interface
type RoleRepoInterface interface {
	Add(role Organization_Roles) Organization_Roles
	Get() []Organization_Roles
	GetByID(id int) Organization_Roles
	Update(role Organization_Roles) Organization_Roles
	Delete(id int) Organization_Roles
	Assign(roleID int, permissions []int) bool
	UnAssign(roleID int, permissions []int) bool
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		db: db,
	}
}

// Add add role
func (roleRepo *RoleRepo) Add(role Organization_Roles) (Organization_Roles, error) {
	res := roleRepo.db.Create(&role)
	if res.Error != nil {
		return Organization_Roles{}, res.Error
	}
	return role, nil
}

// Get get all roles
func (repoRole *RoleRepo) Get(organizationID int) ([]Organization_Roles, error) {
	var roles []Organization_Roles
	res := repoRole.db.Raw("SELECT * FROM organization_roles WHERE organization_id=?", organizationID).Scan(&roles)
	if res.Error != nil {
		return []Organization_Roles{}, res.Error
	}
	return roles, nil
}

// GetByID return single role details
func (repoRole *RoleRepo) GetByID(id int) ([]RoleDetails, error) {
	var roles []RoleDetails
	query := ` SELECT r.name , r.description , pr.permission_id , r.id
	FROM public.organization_roles AS r
	INNER JOIN (SELECT * FROM public.permissions_roles WHERE role_id=?) pr
	ON r.id= pr.role_id`
	res := repoRole.db.Raw(query, id).Scan(&roles)
	if res.Error != nil {
		return []RoleDetails{}, res.Error
	}
	return roles, nil
}

// Update update a role
func (roleRepo *RoleRepo) Update(role Organization_Roles) (Organization_Roles, error) {
	res := roleRepo.db.Save(&role)
	if res.Error != nil {
		return Organization_Roles{}, res.Error
	}
	return role, nil
}

// Delete role
func (roleRepo *RoleRepo) Delete(roleID int) error {
	res := roleRepo.db.Delete(Organization_Roles{
		ID: roleID,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Assign permissions to role
func (roleRepo *RoleRepo) AssignPermission(roleID int, permissions []int) (bool, error) {
	var id int
	for _, perm := range permissions {
		res := roleRepo.db.Raw("INSERT INTO permissions_roles (role_id, permission_id) VALUES (?, ?)", roleID, perm).Scan(&id)
		if res.Error != nil {
			return false, res.Error
		}
	}
	return true, nil
}

// UnAssign remove permissions assigning from role
func (roleRepo *RoleRepo) UnAssignPermission(roleID int, permissions []int) (bool, error) {
	var results interface{}
	for _, perm := range permissions {
		res := roleRepo.db.Raw("DELETE FROM permissions_roles WHERE permission_id=? AND role_id=?", perm, roleID).Scan(&results)
		if res.Error != nil {
			return false, res.Error
		}
	}
	return true, nil
}
