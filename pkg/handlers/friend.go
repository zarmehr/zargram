package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zargram/models"
)

func (h *Handler) AddFriend(c *gin.Context) {
	var friend models.Friend
	if err := c.ShouldBindJSON(&friend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.AddFriend(friend)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusCreated, id)

}
func (h *Handler) GetFriendsByUserID(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	friends, err := h.services.GetFriendsByUserID(userID)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusOK, friends)

}

func (h *Handler) DeleteFriendByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
			return
		}
	}
	h.services.DeleteFriendByID(id)
	c.JSON(http.StatusOK, statusResponse{"friend deleted"})
}
