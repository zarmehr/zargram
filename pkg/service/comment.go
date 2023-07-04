package service

import (
	"zargram/models"
	"zargram/pkg/repository"
)

type CommentService struct {
	repo *repository.Repository
}

func NewCommentService(repo *repository.Repository) *CommentService {
	return &CommentService{repo: repo}
}
func (p *CommentService) CreateComment(comment models.Comment) (int, error) {
	return p.repo.CreateComment(comment)
}
func (p *CommentService) GetCommentsByPostID(postID int) ([]models.Comment, error) {
	return p.repo.GetCommentsByPostID(postID)
}

func (p *CommentService) UpdateCommentByID(id int, t models.Comment) error {
	return p.repo.UpdateCommentByID(id, t)
}

func (p *CommentService) DeleteCommentByID(id int) error {
	return p.repo.DeleteCommentByID(id)
}
