package users

type UserService struct {
	userRepo *UserRepo
}

func NewUserService(userRepo *UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
