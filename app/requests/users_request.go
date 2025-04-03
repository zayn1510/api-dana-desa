package requests

import (
	"apidanadesa/app/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var secretKey []byte

type UserRequestCreate struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserRequestLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r *UserRequestCreate) ToModelUser() *models.User {
	pass, err := hashPassword(r.Password)
	if err != nil {
		log.Fatal(err)
	}
	return &models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: pass,
	}
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
