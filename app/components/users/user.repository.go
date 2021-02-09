package users

import (
	"fmt"
	"log"

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

func (userRepo *UserRepo) Add(user *Organizations_Users) (int, error) {
	log.Println("adding user")
	var id int
	res := userRepo.db.Create(&user)
	// res := roleRepo.db.Raw("INSERT INTO organization_roles (name, description, organization_id) VALUES (?, ?, ?) returning id", role.Name, role.Description, role.OrganizationID).Scan(&id)
	fmt.Println("add", id)
	// fmt.Println("add", res.)
	if res.Error != nil {
		return 0, res.Error
	}
	return id, nil
}

func (userRole *UserRepo) Get(organizationID int) ([]Organizations_Users, error) {
	var users []Organizations_Users
	res := userRole.db.Where("organizations_id", organizationID).Find(&users)
	if res.Error != nil {
		return []Organizations_Users{}, res.Error
	}
	return users, nil
}

func (userRole *UserRepo) GetByID(id int) (Organizations_Users, error) {
	var user Organizations_Users
	userToDelete := Organizations_Users{
		ID: id,
	}
	res := userRole.db.Find(&user, userToDelete)
	fmt.Println(user)
	if res.Error != nil {
		return Organizations_Users{}, res.Error
	}
	return user, nil
}

func (userRepo *UserRepo) Update(user *Organizations_Users) (Organizations_Users, error) {
	res := userRepo.db.Save(&user)
	// res := roleRepo.db.Raw("UPDATE organization_roles SET name=?, description=? WHERE id=?", role.Name, role.Description, role.ID).Scan(&role)
	if res.Error != nil {
		return Organizations_Users{}, res.Error
	}
	return *user, nil
}

// Delete role
func (userRepo *UserRepo) Delete(userID int) error {
	// user :=
	res := userRepo.db.Delete(Organizations_Users{
		ID: userID,
	})
	// res := roleRepo.db.Raw("DELETE FROM organization_roles WHERE id=?", roleID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
