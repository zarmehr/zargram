package service

import (
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"zargram/models"
	"zargram/pkg/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) DeleteUserByID(id int) error {
	return u.repo.DeleteUserByID(id)
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserService) UpdateUserByID(id int, t models.User) error {
	return u.repo.UpdateUserByID(id, t)
}

func (u *UserService) GetUserByID(userID int) (models.User, error) {
	return u.repo.GetUserByID(userID)
}
func (u *UserService) UpdateUserAvatar(userID int, avatar string, file multipart.File) (err error) {

	// Check if the user already has an avatar
	user, err := u.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.HasAvatar() {
		return models.ErrUserAlreadyHasAvatar
	}

	// Create a destination file on the server
	savePath := filepath.Join("uploads", avatar)
	dst, err := os.Create(savePath)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create destination file"})
		return err
	}

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	defer func() {
		errClose := dst.Close()

		if errClose != nil {
			err = errors.Wrap(err, errClose.Error())
		}
	}()

	return u.repo.UpdateUserAvatar(userID, avatar)
}
