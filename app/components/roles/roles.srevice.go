package roles

import (
	"net/http"
	"userManagementApi/app/components/users"
	"userManagementApi/app/responses"
)

type RoleService struct {
	roleRepo    *RoleRepo
	userService *users.UserService
}

func NewRoleService(roleRepo *RoleRepo) *RoleService {
	// usersService := new(users.User)
	var roleService RoleService = RoleService{
		roleRepo: roleRepo,
	}
	return &roleService
}

// add role
func (roleServ *RoleService) Add(role RoleCreateDTO) (Organization_Roles, responses.ErrorData) {
	var err error
	var addedRole Organization_Roles
	formattedRole := Organization_Roles{
		Name:           role.Name,
		OrganizationID: role.OrganizationID,
		Description:    role.Description,
	}

	addedRole, err = roleServ.roleRepo.Add(formattedRole)
	if err != nil {
		return Organization_Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	//check permission length
	if len(role.Permissions) != 0 {
		_, err = roleServ.roleRepo.AssignPermission(addedRole.ID, role.Permissions)
	}
	if err != nil {
		return Organization_Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return addedRole, responses.ErrorData{}
}

// update role
func (roleServ *RoleService) Update(role RoleUpdateDTO) (Organization_Roles, responses.ErrorData) {
	var err error
	if len(role.UnAssign) != 0 {
		_, err = roleServ.roleRepo.UnAssignPermission(role.ID, role.UnAssign)
	}
	if err != nil {
		return Organization_Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	if len(role.NewAssign) != 0 {
		_, err = roleServ.roleRepo.AssignPermission(role.ID, role.NewAssign)
	}
	if err != nil {
		return Organization_Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	var r Organization_Roles = Organization_Roles{
		Name:           role.Name,
		Description:    role.Description,
		ID:             role.ID,
		OrganizationID: role.OrganizationID,
	}
	r, err = roleServ.roleRepo.Update(r)
	if err != nil {
		return Organization_Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return r, responses.ErrorData{}
}

// DeleteRole delete role
func (roleServ *RoleService) DeleteRole(roleID int) responses.ErrorData {
	isRoleUsed, errors := users.Service.IsRoleAssignedToUser(roleID)
	if errors.Errors != nil {
		return errors
	}
	if isRoleUsed {
		return responses.HandleError(http.StatusMethodNotAllowed, "Can't delete role already used")
	}
	err := roleServ.roleRepo.Delete(roleID)
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return responses.ErrorData{}
}

// Get get all roles
func (roleServ *RoleService) Get(organizationID int) ([]Organization_Roles, responses.ErrorData) {
	roles, err := roleServ.roleRepo.Get(organizationID)
	if err != nil {
		return []Organization_Roles{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return roles, responses.ErrorData{}
}

// get role by id
func (roleServ *RoleService) GetRoleByID(roleID int) (SingleRole, responses.ErrorData) {
	// var roleDetails RoleDetails
	roles, err := roleServ.roleRepo.GetByID(roleID)
	if err != nil {
		return SingleRole{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
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
	return roleDetails, responses.ErrorData{}
}
