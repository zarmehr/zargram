package service

import (
	"zargram/models"
	"zargram/pkg/repository"
)

type PostService struct {
	repo *repository.Repository
}

func NewPostService(repo *repository.Repository) *PostService {
	return &PostService{repo: repo}
}
func (p *PostService) CreatePost(post models.Post) (int, error) {
	return p.repo.CreatePost(post)
}
func (p *PostService) GetPostsByUserID(userID int) ([]models.Post, error) {
	return p.repo.GetPostsByUserID(userID)
}

func (p *PostService) UpdatePostByID(id int, t models.Post) error {
	return p.repo.UpdatePostByID(id, t)
}

func (p *PostService) DeletePostByID(id int) error {
	return p.repo.DeletePostByID(id)
}
func (p *PostService) MarkArchived(id int, t models.Post) error {
	return p.repo.MarkArchived(id, t)
}
func (p *PostService) GetPostByID(id int) (t models.Post, err error) {
	return p.repo.GetPostByID(id)
}
