package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zargram/models"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.CreatePost(post)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusNotFound, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusCreated, id)

}
func (h *Handler) GetPostsByUserID(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	posts, err := h.services.GetPostsByUserID(userID)
	if err != nil {
		if errors.Is(err, models.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}
	c.JSON(http.StatusOK, posts)

}

func (h *Handler) UpdatePostByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if errors.Is(err, models.ErrInvalidId) {
		c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
		return
	}

	var t *models.Post
	if err = c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}

	if err = h.services.UpdatePostByID(id, *t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "error while updating post",
		})
		return
	}

	c.JSON(http.StatusOK, statusResponse{"successfully updated"})
}

func (h *Handler) DeletePostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
			return
		}
	}
	h.services.DeletePostByID(id)
	c.JSON(http.StatusOK, statusResponse{"post deleted"})
}

func (h *Handler) MarkArchived(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if errors.Is(err, models.ErrInvalidId) {
		c.JSON(http.StatusNotFound, models.ErrInvalidId.Error())
		return
	}
	var t models.Post
	t, err = h.services.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	if err = h.services.MarkArchived(id, t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "error while updating user",
		})
		return
	}

	c.JSON(http.StatusOK, statusResponse{"successfully updated"})
}
