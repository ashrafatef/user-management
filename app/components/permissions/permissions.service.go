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

func (permServ *PermissionService) Create(permission *Permissions) {
	// CreatePermissionRepo(permission)
	permServ.permissionRepo.CreatePermissionRepo(permission)
}

func (permServ *PermissionService) GetAllPermissions() *[]Permissions {
	// CreatePermissionRepo(permission)
	return permServ.permissionRepo.GetAll()
}

func (permServ *PermissionService) UpdatePermission(permission *Permissions) {
	// CreatePermissionRepo(permission)
	permServ.permissionRepo.UpdatePermissionRepo(permission)
}
