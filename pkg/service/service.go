package service

import (
	"mime/multipart"
	"zargram/models"
	"zargram/pkg/repository"
)

type Authorization interface {
	CreateUser(u models.User) (int, error)
	GetUser(username, password string) (models.User, error)
	GenerateToken(username, password string) (string, error)
}
type User interface {
	GetAllUsers() ([]models.User, error)
	UpdateUserByID(id int, t models.User) error
	DeleteUserByID(id int) error
	GetUserByID(userID int) (models.User, error)
	UpdateUserAvatar(userID int, filename string, file multipart.File) error
}
type Post interface {
	CreatePost(post models.Post) (int, error)
	GetPostsByUserID(userID int) ([]models.Post, error)
	UpdatePostByID(id int, t models.Post) error
	DeletePostByID(id int) error
	MarkArchived(id int, t models.Post) error
	GetPostByID(id int) (t models.Post, err error)
}
type Comment interface {
	CreateComment(comment models.Comment) (int, error)
	GetCommentsByPostID(postID int) ([]models.Comment, error)
	UpdateCommentByID(id int, t models.Comment) error
	DeleteCommentByID(id int) error
}
type Story interface {
	CreateStory(story models.Story) (int, error)
	GetStoriesByUserID(userID int) ([]models.Story, error)
	UpdateStoryByID(id int, t models.Story) error
	DeleteStoryByID(id int) error
}

type Friend interface {
	AddFriend(friend models.Friend) (int, error)
	GetFriendsByUserID(userID int) ([]models.Friend, error)
	DeleteFriendByID(id int) error
}
type Reaction interface {
	CreateReaction(reaction models.Reaction) (int, error)
	GetReactionsByPostID(postID int) ([]models.Reaction, error)
	UpdateReactionByID(id int, t models.Reaction) error
	DeleteReactionByID(id int) error
}

type Service struct {
	Authorization
	User
	Post
	Comment
	Story
	Friend
	Reaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		User:          NewUserService(repos),
		Post:          NewPostService(repos),
		Comment:       NewCommentService(repos),
		Story:         NewStoryService(repos),
		Friend:        NewFriendService(repos),
		Reaction:      NewReactionService(repos),
	}
}
