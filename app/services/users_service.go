package services

import (
	"apidanadesa/app/middleware"
	"apidanadesa/app/models"
	"apidanadesa/app/requests"
	"apidanadesa/config"
	"errors"
	"gorm.io/gorm"
)

type UsersService struct {
	db *gorm.DB
}

func NewUsersService() *UsersService {
	return &UsersService{
		db: config.GetDB(),
	}
}

func (s *UsersService) CreateUser(req requests.UserRequestCreate) error {
	data := req.ToModelUser()
	err := s.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *UsersService) LoginUser(req requests.UserRequestLogin) (string, error) {
	var user models.User
	result := s.db.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return "", errors.New("user not found")
	}

	if err := requests.CheckPassword(user.Password, req.Password); err != nil {
		return "", errors.New("invalid credentials")
	}
	return middleware.GenerateJWT(user.Username)
}
