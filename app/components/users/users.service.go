package users

import (
	"errors"
	"net/http"
	"userManagementApi/app/responses"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *UserRepo
}

func NewUserService(userRepo *UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Get get all roles
func (userService *UserService) Get(organizationID int) ([]Organizations_Users, responses.ErrorData) {
	users, err := userService.userRepo.Get(organizationID)
	if err != nil {
		return []Organizations_Users{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return users, responses.ErrorData{}
}

func (userService *UserService) GetUserByID(userID int) (Organizations_Users, responses.ErrorData) {
	user, err := userService.userRepo.GetByID(userID)
	if err != nil {
		return Organizations_Users{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return user, responses.ErrorData{}
}

func (userService *UserService) IsRoleAssignedToUser(roleID int) (bool, responses.ErrorData) {
	user, err := userService.userRepo.GetFirstByRole(roleID)
	isErrorRecordNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if isErrorRecordNotFound && user.ID == 0 {
		return false, responses.ErrorData{}
	}
	if !isErrorRecordNotFound && err != nil {
		return false, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return true, responses.ErrorData{}
}

// DeleteRole delete role
func (userService *UserService) Delete(userID int) responses.ErrorData {
	err := userService.userRepo.Delete(userID)
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return responses.ErrorData{}
}

// add role
func (userService *UserService) Add(user UserCreateDTO) (Organizations_Users, responses.ErrorData) {
	// var err error
	// var id int
	u := Organizations_Users{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Password:        user.Password,
		Salt:            user.Salt,
		OrganizationsID: user.OrganizationID,
		RoleID:          user.RoleID,
	}
	u, err := userService.userRepo.Add(&u)
	if err != nil {
		return Organizations_Users{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return u, responses.ErrorData{}
}

// Update update user
func (userService *UserService) Update(user UserUpdateDTO) (Organizations_Users, responses.ErrorData) {
	var err error
	// do update role attributes
	u := Organizations_Users{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Salt:            user.Salt,
		OrganizationsID: user.OrganizationID,
		RoleID:          user.RoleID,
	}
	u, err = userService.userRepo.Update(&u)
	if err != nil {
		return Organizations_Users{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return u, responses.ErrorData{}
}
