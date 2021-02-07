package roles

import (
	"fmt"
	"net/http"
	error_handler "userManagementApi/app/error_handler"
)

type RoleService struct {
	roleRepo *RoleRepo
}

func NewRoleService(roleRepo *RoleRepo) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

// add role
func (roleServ *RoleService) Add(role RoleCreateDTO) (Roles, error) {
	var err error
	var id int
	r := Roles{
		Name:           role.Name,
		OrganizationID: role.OrganizationID,
		Description:    role.Description,
	}
	id, err = roleServ.roleRepo.Add(r)
	fmt.Println("add service", id)
	if err != nil {
		return Roles{}, error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	//check permission length
	if len(role.Permissions) != 0 {
		_, err = roleServ.roleRepo.Assign(id, role.Permissions)
	}
	if err != nil {
		return Roles{}, error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	return r, nil
}

// update role
func (roleServ *RoleService) Update(role RoleUpdateDTO) (Roles, error) {
	var err error

	//check unassign array
	if len(role.UnAssign) != 0 {
		// do un assign
		_, err = roleServ.roleRepo.UnAssign(role.ID, role.UnAssign)
	}
	if err != nil {
		return Roles{}, error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	//check assign array
	if len(role.NewAssign) != 0 {
		//do new assign
		_, err = roleServ.roleRepo.Assign(role.ID, role.NewAssign)
	}
	if err != nil {
		return Roles{}, error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	// do update role attributes
	var r Roles = Roles{
		Name:        role.Name,
		Description: role.Description,
		ID:          role.ID,
	}
	r, err = roleServ.roleRepo.Update(r)
	if err != nil {
		return Roles{}, error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	return r, nil
}

// DeleteRole delete role
func (roleServ *RoleService) DeleteRole(roleID int) error {
	err := roleServ.roleRepo.Delete(roleID)
	if err != nil {
		return error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Get get all roles
func (roleServ *RoleService) Get(organizationID int) ([]Roles, error) {
	roles, err := roleServ.roleRepo.Get(organizationID)
	if err != nil {
		return []Roles{}, error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	return roles, nil
}

// get role by id
func (roleServ *RoleService) GetRoleByID(roleID int) (SingleRole, error) {
	// var roleDetails RoleDetails
	roles, err := roleServ.roleRepo.GetByID(roleID)
	if len(roles) < 1 {
		return SingleRole{}, nil
	}
	var roleDetails = SingleRole{
		ID:          roles[0].ID,
		Name:        roles[0].Name,
		Description: roles[0].Description,
	}
	for _, role := range roles {
		roleDetails.Permissions = append(roleDetails.Permissions, role.PermissionID)
	}
	fmt.Println(roleDetails)
	if err != nil {
		return SingleRole{}, error_handler.HandleError(http.StatusInternalServerError, err.Error())
	}
	return roleDetails, nil
}
