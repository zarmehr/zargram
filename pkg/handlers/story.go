package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zargram/models"
)

func (h *Handler) CreateStory(c *gin.Context) {
	var story models.Story
	if err := c.ShouldBindJSON(&story); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.CreateStory(story)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusCreated, id)

}
func (h *Handler) GetStoriesByUserID(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	stories, err := h.services.GetStoriesByUserID(userID)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusOK, stories)

}

func (h *Handler) UpdateStory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if errors.Is(err, models.ErrInvalidId) {
		c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
		return
	}

	var t *models.Story
	if err = c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}

	if err = h.services.UpdateStoryByID(id, *t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "error while updating story",
		})
		return
	}

	c.JSON(http.StatusOK, statusResponse{"successfully updated"})
}

func (h *Handler) DeleteStory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
			return
		}
	}
	h.services.DeleteStoryByID(id)
	c.JSON(http.StatusOK, statusResponse{"story deleted"})
}
