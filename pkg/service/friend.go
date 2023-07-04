package service

import (
	"zargram/models"
	"zargram/pkg/repository"
)

type FriendService struct {
	repo *repository.Repository
}

func NewFriendService(repo *repository.Repository) *FriendService {
	return &FriendService{repo: repo}
}
func (p *FriendService) AddFriend(friend models.Friend) (int, error) {
	return p.repo.AddFriend(friend)
}
func (p *FriendService) GetFriendsByUserID(userID int) ([]models.Friend, error) {
	return p.repo.GetFriendsByUserID(userID)
}

func (p *FriendService) DeleteFriendByID(id int) error {
	return p.repo.DeleteFriendByID(id)
}
