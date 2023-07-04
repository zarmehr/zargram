package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zargram/models"
)

func (h *Handler) CreateReaction(c *gin.Context) {
	var reaction models.Reaction
	if err := c.ShouldBindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.CreateReaction(reaction)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusCreated, id)

}
func (h *Handler) CountReactionsByPostId(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	reactions, err := h.services.GetReactionsByPostID(postID)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusOK, len(reactions))
}

func (h *Handler) UpdateReactionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if errors.Is(err, models.ErrInvalidId) {
		c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
		return
	}

	var t *models.Reaction
	if err = c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}

	if err = h.services.UpdateReactionByID(id, *t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "error while updating reaction",
		})
		return
	}

	c.JSON(http.StatusOK, statusResponse{"successfully updated"})

}

func (h *Handler) DeleteReactionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
			return
		}
	}
	h.services.DeleteReactionByID(id)
	c.JSON(http.StatusOK, statusResponse{"reaction deleted"})
}
