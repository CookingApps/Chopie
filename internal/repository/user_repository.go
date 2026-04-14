package repository

import (
	"Chopie/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

//	func HashPassword(password string) (string, error) {
//		// bcrypt.DefaultCost is 10; higher costs are more secure but slower
//		bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//		return string(bytes), err
//	}
func (r *userRepository) Create(user *model.User) error {
	// hashedPassword, err := HashPassword(user.Password) // Hash the password before saving
	// if err != nil {
	// 	return err
	// }
	// user.Password = hashedPassword
	return r.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Printf("Error finding user by email: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}
