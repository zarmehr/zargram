package repository

import (
	"gorm.io/gorm"
	"zargram/models"
)

type Authorization interface {
	CreateUser(u models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}
type User interface {
	GetAllUsers() (users []models.User, err error)
	UpdateUserByID(id int, t models.User) error
	DeleteUserByID(id int) (err error)
	GetUserByID(userID int) (models.User, error)
	UpdateUserAvatar(userID int, avatar string) error
}
type Post interface {
	CreatePost(post models.Post) (id int, err error)
	GetPostsByUserID(userID int) (posts []models.Post, err error)
	UpdatePostByID(id int, t models.Post) error
	DeletePostByID(id int) (err error)
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

type Repository struct {
	Authorization
	User
	Post
	Comment
	Story
	Friend
	Reaction
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Post:          NewPostPostgres(db),
		Comment:       NewCommentPostgres(db),
		Story:         NewStoryPostgres(db),
		Friend:        NewFriendPostgres(db),
		Reaction:      NewReactionPostgres(db),
	}
}
