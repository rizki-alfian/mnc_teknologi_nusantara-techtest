package providers

import (
	"mnc-users/apps/databases/repositories"
	"mnc-users/apps/users"
)

func NewUserService(userRepo *repositories.UserRepository) *users.UserService {
	return &users.UserService{Repo: userRepo}
}
