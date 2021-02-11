package roles

import (
	"fmt"

	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

// RoleRepoInterface role repo interface
type RoleRepoInterface interface {
	Add(role Organizations_Roles) Organizations_Roles
	Get() []Organizations_Roles
	GetByID(id int) Organizations_Roles
	Update(role Organizations_Roles) Organizations_Roles
	Delete(id int) Organizations_Roles
	Assign(roleID int, permissions []int) bool
	UnAssign(roleID int, permissions []int) bool
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		db: db,
	}
}

// Add add role
func (roleRepo *RoleRepo) Add(role Organizations_Roles) (Organizations_Roles, error) {
	res := roleRepo.db.Create(&role)
	// res := roleRepo.db.Raw("INSERT INTO organization_roles (name, description, organization_id) VALUES (?, ?, ?) returning id", role.Name, role.Description, role.OrganizationID).Scan(&id)
	if res.Error != nil {
		return Organizations_Roles{}, res.Error
	}
	return role, nil
}

// Get get all roles
func (repoRole *RoleRepo) Get(organizationID int) ([]Organizations_Roles, error) {
	var roles []Organizations_Roles
	res := repoRole.db.Raw("SELECT * FROM organization_roles WHERE organization_id=?", organizationID).Scan(&roles)
	if res.Error != nil {
		return []Organizations_Roles{}, res.Error
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
func (roleRepo *RoleRepo) Update(role Organizations_Roles) (Organizations_Roles, error) {
	// res := roleRepo.db.Save(&role)
	res := roleRepo.db.Save(&role)
	// res := roleRepo.db.Raw("UPDATE organization_roles SET name=?, description=? WHERE id=?", role.Name, role.Description, role.ID).Scan(&role)
	fmt.Println(role)
	if res.Error != nil {
		return Organizations_Roles{}, res.Error
	}
	return role, nil
}

// Delete role
func (roleRepo *RoleRepo) Delete(roleID int) error {

	res := roleRepo.db.Delete(Organizations_Roles{
		ID: roleID,
	})
	// res := roleRepo.db.Raw("DELETE FROM organization_roles WHERE id=?", roleID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Assign permissions to role
func (roleRepo *RoleRepo) Assign(roleID int, permissions []int) (bool, error) {
	var id int
	fmt.Println("assign", roleID)
	for _, perm := range permissions {
		res := roleRepo.db.Raw("INSERT INTO permissions_roles (role_id, permission_id) VALUES (?, ?)", roleID, perm).Scan(&id)
		if res.Error != nil {
			return false, res.Error
		}
	}
	return true, nil
}

// UnAssign remove permissions assigning from role
func (roleRepo *RoleRepo) UnAssign(roleID int, permissions []int) (bool, error) {
	for _, perm := range permissions {
		res := roleRepo.db.Raw("DELETE * FROM permissions_roles WHERE permission_id=? AND role_id=?", roleID, perm)
		if res.Error != nil {
			return false, res.Error
		}
	}
	return true, nil
}
