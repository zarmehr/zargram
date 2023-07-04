package handlers

import (
	"net/http"
	"strconv"
	"zargram/models"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	t, err := h.services.GetAllUsers()
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusOK, t)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	user, err := h.services.GetUserByID(id)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.CreateUser(user)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusCreated, id)

}

func (h *Handler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if errors.Is(err, models.ErrInvalidId) {
		c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
		return
	}

	var t *models.User
	if err = c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}

	if err = h.services.UpdateUserByID(id, *t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "error while updating user",
		})
		return
	}

	c.JSON(http.StatusOK, statusResponse{"successfully updated"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
			return
		}
	}
	h.services.DeleteUserByID(id)
	c.JSON(http.StatusOK, statusResponse{"user deleted"})
}

//=============================================================

const (
	uploadAvatarFormKey = "file"
)

func (h *Handler) uploadAvatarHandler(c *gin.Context) {
	// Retrieve the uploaded file
	file, header, err := c.Request.FormFile(uploadAvatarFormKey)
	if err != nil {
		if errors.Is(err, models.ErrFailToRetrieve) {
			c.JSON(http.StatusNotFound, models.ErrFailToRetrieve.Error())
			return
		}
	}

	defer file.Close()

	// Retrieve the user ID from the request
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
			return
		}

	}

	filename := header.Filename

	// Update the user's avatar in the database
	err = h.services.UpdateUserAvatar(userId, filename, file)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, models.ErrUserNotFound.Error())
			return
		}

		if errors.Is(err, models.ErrUserAlreadyHasAvatar) {
			c.JSON(http.StatusBadRequest, models.ErrUserAlreadyHasAvatar.Error())
			return
		}

		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}

	c.JSON(http.StatusOK, statusResponse{"File uploaded successfully"})
}

type UploadAvatarResponse struct {
}
