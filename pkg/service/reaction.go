package service

import (
	"zargram/models"
	"zargram/pkg/repository"
)

type ReactionService struct {
	repo *repository.Repository
}

func NewReactionService(repo *repository.Repository) *ReactionService {
	return &ReactionService{repo: repo}
}

func (p *ReactionService) CreateReaction(reaction models.Reaction) (int, error) {
	return p.repo.CreateReaction(reaction)

}
func (p *ReactionService) GetReactionsByPostID(postID int) ([]models.Reaction, error) {
	return p.repo.GetReactionsByPostID(postID)
}

func (p *ReactionService) UpdateReactionByID(id int, t models.Reaction) error {
	return p.repo.UpdateReactionByID(id, t)
}

func (p *ReactionService) DeleteReactionByID(id int) error {
	return p.repo.DeleteReactionByID(id)
}
