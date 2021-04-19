package services

import (
	"go-fiber-auth-test/db"
	"go-fiber-auth-test/domain"
	"go-fiber-auth-test/dto"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(request dto.CreateUserRequest) (*domain.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Username: request.Username,
		Password: password,
	}

	result := db.GetConnectionDB().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := db.GetConnectionDB().Where(domain.User{
		Username: username,
	}).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
