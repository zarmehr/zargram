package service

import (
	"zargram/models"
	"zargram/pkg/repository"
)

type StoryService struct {
	repo *repository.Repository
}

func NewStoryService(repo *repository.Repository) *StoryService {
	return &StoryService{repo: repo}
}

func (s *StoryService) CreateStory(story models.Story) (int, error) {
	return s.repo.CreateStory(story)
}
func (s *StoryService) GetStoriesByUserID(userID int) ([]models.Story, error) {
	return s.repo.GetStoriesByUserID(userID)
}

func (s *StoryService) UpdateStoryByID(id int, t models.Story) error {
	return s.repo.UpdateStoryByID(id, t)
}

func (s *StoryService) DeleteStoryByID(id int) error {
	return s.repo.DeleteStoryByID(id)
}
