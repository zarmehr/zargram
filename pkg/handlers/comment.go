package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zargram/models"
)

func (h *Handler) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.CreateComment(comment)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusCreated, id)

}
func (h *Handler) GetCommentsByPostId(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	comments, err := h.services.GetCommentsByPostID(postID)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusOK, comments)
}

func (h *Handler) UpdateCommentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if errors.Is(err, models.ErrInvalidId) {
		c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
		return
	}

	var t *models.Comment
	if err = c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}

	if err = h.services.UpdateCommentByID(id, *t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "error while updating comment",
		})
		return
	}

	c.JSON(http.StatusOK, statusResponse{"successfully updated"})

}

func (h *Handler) DeleteCommentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
			return
		}
	}
	h.services.DeleteCommentByID(id)
	c.JSON(http.StatusOK, statusResponse{"comment deleted"})
}
