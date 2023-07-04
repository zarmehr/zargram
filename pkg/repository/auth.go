package repository

import (
	"zargram/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(u models.User) (id int, err error) {
	user := models.User{FullName: u.FullName, Username: u.Username, Password: u.Password}
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *AuthPostgres) GetUser(username, password string) (u models.User, err error) {

	result := r.db.Debug().Take(&u, "username=? AND password=?", username, password)
	if result.Error != nil {
		return models.User{}, err
	}

	return u, nil

}
