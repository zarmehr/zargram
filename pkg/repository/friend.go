package repository

import (
	"gorm.io/gorm"
	"zargram/models"
)

type FriendPostgres struct {
	db *gorm.DB
}

func NewFriendPostgres(db *gorm.DB) *FriendPostgres {
	return &FriendPostgres{db: db}
}
func (p *FriendPostgres) AddFriend(friend models.Friend) (id int, err error) {
	friend = models.Friend{UserID: friend.UserID, FriendID: friend.FriendID}
	err = p.db.Debug().Create(&friend).Error
	if err != nil {
		return 0, err
	}
	return friend.ID, nil

}

func (p *FriendPostgres) GetFriendsByUserID(userID int) (friends []models.Friend, err error) {
	err = p.db.Where("user_id = ?", userID).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (p *FriendPostgres) DeleteFriendByID(id int) (err error) {
	err = p.db.Delete(&models.Friend{}, id).Error
	if err != nil {
		panic(err)
	}
	return nil
}
