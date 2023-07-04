package handlers

import (
	"net/http"
	"zargram/pkg/service"

	"github.com/gin-gonic/gin"
)

//------------------------------------

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//-------------------------------------

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/", Ping)

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}

	api := router.Group("/api")

	users := api.Group("/users", h.userIdentity)
	{
		users.GET("/", h.GetAllUsers)
		users.GET("/:id", h.GetUserByID)
		//users.POST("/", h.CreateUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
		users.POST("/avatar/:id", h.uploadAvatarHandler)
	}
	comments := api.Group("/comments", h.userIdentity)
	{
		comments.GET("/:id", h.GetCommentsByPostId)
		comments.POST("/", h.CreateComment)
		comments.PUT("/:id", h.UpdateCommentByID)
		comments.DELETE("/:id", h.DeleteCommentByID)
	}
	posts := api.Group("/posts", h.userIdentity)
	{
		posts.GET("/", h.GetPostsByUserID)
		posts.POST("/", h.CreatePost)
		posts.PUT("/:id", h.UpdatePostByID)
		posts.DELETE("/:id", h.DeletePostByID)
		archived := posts.Group("/archived", h.userIdentity)
		archived.PUT("/:id", h.MarkArchived)
	}
	stories := api.Group("/stories", h.userIdentity)
	{
		stories.GET("/", h.GetStoriesByUserID)
		stories.POST("/", h.CreateStory)
		stories.PUT("/:id", h.UpdateStory)
		stories.DELETE("/:id", h.DeleteStory)
	}
	followers := api.Group("/followers", h.userIdentity)
	{
		followers.GET("/", h.GetFriendsByUserID)
		followers.POST("/", h.AddFriend)
		followers.DELETE("/:id", h.DeleteFriendByID)
	}
	reactions := api.Group("/reactions", h.userIdentity)
	{
		reactions.GET("/:id", h.CountReactionsByPostId)
		reactions.POST("/", h.CreateReaction)
		reactions.PUT("/:id", h.UpdateReactionByID)
		reactions.DELETE("/:id", h.DeleteReactionByID)
	}

	return router
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"reason": "up and working",
	})
}
