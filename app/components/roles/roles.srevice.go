package roles

import (
	"fmt"
	"net/http"
	"userManagementApi/app/responses"
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
func (roleServ *RoleService) Add(role RoleCreateDTO) (Roles, responses.ErrorData) {
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
		return Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	//check permission length
	if len(role.Permissions) != 0 {
		_, err = roleServ.roleRepo.Assign(id, role.Permissions)
	}
	if err != nil {
		return Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return r, responses.ErrorData{}
}

// update role
func (roleServ *RoleService) Update(role RoleUpdateDTO) (Roles, responses.ErrorData) {
	var err error

	//check unassign array
	if len(role.UnAssign) != 0 {
		// do un assign
		_, err = roleServ.roleRepo.UnAssign(role.ID, role.UnAssign)
	}
	if err != nil {
		return Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	//check assign array
	if len(role.NewAssign) != 0 {
		//do new assign
		_, err = roleServ.roleRepo.Assign(role.ID, role.NewAssign)
	}
	if err != nil {
		return Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	// do update role attributes
	var r Roles = Roles{
		Name:        role.Name,
		Description: role.Description,
		ID:          role.ID,
	}
	r, err = roleServ.roleRepo.Update(r)
	if err != nil {
		return Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return r, responses.ErrorData{}
}

// DeleteRole delete role
func (roleServ *RoleService) DeleteRole(roleID int) responses.ErrorData {
	err := roleServ.roleRepo.Delete(roleID)
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return responses.ErrorData{}
}

// Get get all roles
func (roleServ *RoleService) Get(organizationID int) ([]Roles, responses.ErrorData) {
	roles, err := roleServ.roleRepo.Get(organizationID)
	if err != nil {
		return []Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return roles, responses.ErrorData{}
}

// get role by id
func (roleServ *RoleService) GetRoleByID(roleID int) (SingleRole, responses.ErrorData) {
	// var roleDetails RoleDetails
	roles, err := roleServ.roleRepo.GetByID(roleID)
	if len(roles) < 1 {
		return SingleRole{}, responses.ErrorData{}
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
		return SingleRole{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return roleDetails, responses.ErrorData{}
}
