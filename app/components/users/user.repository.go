package users

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (userRepo *UserRepo) Add(user *Organizations_Users) (Organizations_Users, error) {
	res := userRepo.db.Create(&user)
	if res.Error != nil {
		return Organizations_Users{}, res.Error
	}
	return *user, nil
}

func (userRole *UserRepo) Get(organizationID int) ([]Organizations_Users, error) {
	var users []Organizations_Users
	res := userRole.db.Where("organizations_id", organizationID).Find(&users)
	if res.Error != nil {
		return []Organizations_Users{}, res.Error
	}
	return users, nil
}

func (userRole *UserRepo) GetFirstByRole(roleID int) (Organizations_Users, error) {
	var user Organizations_Users
	res := userRole.db.First(&user, "role_id=?", roleID)
	if res.Error != nil {
		return Organizations_Users{}, res.Error
	}
	return user, nil
}

func (userRole *UserRepo) GetByID(id int) (Organizations_Users, error) {
	var user Organizations_Users
	userToDelete := Organizations_Users{
		ID: id,
	}
	res := userRole.db.Find(&user, userToDelete)
	if res.Error != nil {
		return Organizations_Users{}, res.Error
	}
	return user, nil
}

func (userRepo *UserRepo) Update(user *Organizations_Users) (Organizations_Users, error) {
	res := userRepo.db.Updates(&user)
	if res.Error != nil {
		return Organizations_Users{}, res.Error
	}
	return *user, nil
}

// Delete role
func (userRepo *UserRepo) Delete(userID int) error {
	res := userRepo.db.Delete(Organizations_Users{
		ID: userID,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
