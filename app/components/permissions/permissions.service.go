package permissions

type PermissionService struct {
	permissionRepo *PermissionRepo
}

func NewPermissionService(permRepo *PermissionRepo) *PermissionService {
	return &PermissionService{
		permissionRepo: permRepo,
	}
}

// func (permServ *PermissionService) GetAllPermissions() {

// }

func (permServ *PermissionService) CreatePermission(permission *PermissionsCreateDTO) {
	// CreatePermissionRepo(permission)
	perm := Permissions{
		CategoryID: permission.CategoryID,
		Type:       permission.Type,
		Name:       permission.Name,
	}
	permServ.permissionRepo.CreatePermissionRepo(&perm)
}

func (permServ *PermissionService) GetAllPermissions() *[]Permissions {
	// CreatePermissionRepo(permission)
	return permServ.permissionRepo.GetAll()
}

func (permServ *PermissionService) UpdatePermission(permission *PermissionsUpdateDTO) {
	// CreatePermissionRepo(permission)
	perm := Permissions{
		ID:   permission.ID,
		Name: permission.Name,
	}
	permServ.permissionRepo.UpdatePermissionRepo(&perm)
}
