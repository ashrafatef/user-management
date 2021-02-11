package permissions

import (
	"net/http"
	"userManagementApi/app/responses"
)

type PermissionService struct {
	permissionRepo *PermissionRepo
}

func NewPermissionService(permRepo *PermissionRepo) *PermissionService {
	return &PermissionService{
		permissionRepo: permRepo,
	}
}

func (permServ *PermissionService) CreatePermission(permission *PermissionsCreateDTO) responses.ErrorData {
	// CreatePermissionRepo(permission)
	perm := Permissions{
		Category: permission.Category,
		Type:     permission.Type,
		Name:     permission.Name,
	}
	err := permServ.permissionRepo.CreatePermissionRepo(&perm)
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return responses.ErrorData{}
}

func (permServ *PermissionService) GetAllPermissions() ([]Permissions, responses.ErrorData) {
	// CreatePermissionRepo(permission)

	permissions, err := permServ.permissionRepo.GetAllPermissionsRepo()
	if err != nil {
		return []Permissions{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return permissions, responses.ErrorData{}
}

func (permServ *PermissionService) UpdatePermission(permission *PermissionsUpdateDTO) responses.ErrorData {
	perm := Permissions{
		ID:   permission.ID,
		Name: permission.Name,
	}
	err := permServ.permissionRepo.UpdatePermissionRepo(&perm)
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return responses.ErrorData{}
}

func (permServ *PermissionService) DeletePermission(permID int) responses.ErrorData {
	var err error
	var permissionID int64
	permissionID, err = permServ.permissionRepo.CheckPermissionAssigning(permID)
	if permissionID != 0 {
		return responses.HandleError(http.StatusMethodNotAllowed, "")
	}
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	err = permServ.permissionRepo.DeletePermissionRepo(permID)
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return responses.ErrorData{}
}
