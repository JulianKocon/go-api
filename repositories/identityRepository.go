package repositories

import (
	"example/go-api/initializers"
	"example/go-api/models"
)

type IdentityRepostiory interface {
	Register(models.User) error
	GetUserByEmail(string) (*models.User, error)
}

type identityRepostiory struct {
}

func NewIdentityRepository() IdentityRepostiory {
	return identityRepostiory{}
}

func (identityRepostiory) Register(user models.User) error{
	if err := user.HashPassword(user.Password); err != nil{
		return err
	}
	initializers.DB.Create(&user)
	return nil
}

func (identityRepostiory) GetUserByEmail(email string) (*models.User, error){
	var user models.User
	// check if email exists and password is correct
	record := initializers.DB.Where("email = ?", email).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}
	return &user, nil
}
