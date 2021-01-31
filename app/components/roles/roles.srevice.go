package roles

import (
	"log"
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
	log.Println("adding user")
	r := Roles{
		Name:           role.Name,
		OrganizationID: role.OrganizationID,
		Description:    role.Description,
	}
	id, err := roleServ.roleRepo.Add(r)
	if err != nil {
		// e := error_handler.FormateErrorResponse(http.StatusInternalServerError, err)
		return Roles{}, error_handler.HandleError(http.StatusInternalServerError, "")
	}
	// id := roleServ.roleRepo.Add(r)

	//check permission length
	if len(role.Permissions) != 0 {
		roleServ.roleRepo.Assign(id, role.Permissions)
	}
	return r, nil
}

// update role
func (roleServ *RoleService) Update(role RoleUpdateDTO) Roles {
	//check unassign array
	if len(role.UnAssign) != 0 {
		// do un assign
		roleServ.roleRepo.UnAssign(role.ID, role.UnAssign)
	}
	//check assign array
	if len(role.NewAssign) != 0 {
		//do new assign
		roleServ.roleRepo.Assign(role.ID, role.NewAssign)
	}
	// do update role attributes
	var r Roles = Roles{
		Name:        role.Name,
		Description: role.Description,
		ID:          role.ID,
	}
	roleServ.roleRepo.Update(r)
	return Roles{}
}

// assign permission
// func (roleServ *RoleService) Assign
// unassign permission

// delete role

// get all roles
// Get get all roles
func (roleServ *RoleService) Get() []Roles {
	return roleServ.roleRepo.Get()
}

// get role by id
