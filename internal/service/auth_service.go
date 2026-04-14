package service

import (
	"Chopie/internal/model"
	"Chopie/internal/repository"
	"Chopie/internal/util"
	"fmt"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func (a *AuthService) Register(email string, password string) (err error) {

	user, err := a.userRepo.FindByEmail(email)

	if err != nil {
		hashedPwd, err := util.HashPassword(password)
		if err != nil {
			return fmt.Errorf("Problem with password ")
		}

		user = &model.User{
			Email:    email,
			Password: hashedPwd,
			// F
		}

		err = a.userRepo.Create(user)
		if err != nil {
			return fmt.Errorf("i no fit create user: %w", err)
		}

		return err
	}

	return fmt.Errorf("User don get account o na to chop remain!")
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}
