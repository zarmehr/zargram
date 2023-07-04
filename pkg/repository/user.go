package repository

import (
	"errors"
	"zargram/models"

	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAllUsers() (users []models.User, err error) {
	err = r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserPostgres) UpdateUserByID(id int, t models.User) error {
	t.ID = id
	tx := r.db.Where("id = ?", id).Save(t)

	return tx.Error
}

func (r *UserPostgres) DeleteUserByID(id int) (err error) {
	err = r.db.Delete(&models.User{}, id).Error
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *UserPostgres) GetUserByID(userID int) (models.User, error) {
	var user models.User
	result := r.db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, models.ErrUserNotFound
		}

		return models.User{}, result.Error
	}
	return user, nil
}

func (r *UserPostgres) UpdateUserAvatar(userID int, avatar string) error {
	var user models.User
	err := r.db.Debug().First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrUserNotFound
		}
		return err
	}

	user.Avatar = avatar
	result := r.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
