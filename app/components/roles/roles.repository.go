package roles

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

// RoleRepoInterface role repo interface
type RoleRepoInterface interface {
	Add(role Roles) Roles
	Get() []Roles
	GetByID(id int) Roles
	Update(role Roles) Roles
	Delete(id int) Roles
	Assign(roleID int, permissions []int) bool
	UnAssign(roleID int, permissions []int) bool
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		db: db,
	}
}

// Add add role
func (roleRepo *RoleRepo) Add(role Roles) (int, error) {
	log.Println("adding user")
	var id int
	err := roleRepo.db.Raw("INSERT INTO oorganization_roles (name, description, organization_id) VALUES (?, ?, ?) returning id", role.Name, role.Description, role.OrganizationID).Scan(&id)
	log.Println(id)
	log.Println("teeeeeeeeeeeeeeeeeeeeeest", err.Error.Error())
	if err != nil {
		return 0, err.Error
	}
	// return id, nil
	return id, nil
}

// Get get all roles
func (repoRole *RoleRepo) Get() []Roles {
	var roles []Roles
	repoRole.db.Raw("SELECT * FROM organization_roles").Scan(&roles)
	return roles
}

// GetByID return single role details
func (repoRole *RoleRepo) GetByID(id int) Roles {

	return Roles{}
}

// Update update a role
func (roleRepo *RoleRepo) Update(role Roles) Roles {
	roleRepo.db.Save(&role)
	return role
}

// Delete role
func (roleRepo *RoleRepo) Delete(id int) Roles {
	return Roles{}
}

// Assign permissions to role
func (roleRepo *RoleRepo) Assign(roleID int, permissions []int) bool {
	var id int
	fmt.Println(roleID)
	for _, perm := range permissions {
		roleRepo.db.Raw("INSERT INTO permissions_roles (role_id, permission_id) VALUES (?, ?)", roleID, perm).Scan(&id)
		fmt.Println(perm)
	}
	return true
}

// UnAssign remove permissions assigning from role
func (roleRepo *RoleRepo) UnAssign(roleID int, permissions []int) bool {
	for _, perm := range permissions {
		roleRepo.db.Raw("DELETE * FROM permissions_roles WHERE permission_id=? AND role_id=?", roleID, perm)
		fmt.Println(perm)
	}
	return true
}
