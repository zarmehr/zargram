package repository

import (
	"gorm.io/gorm"
	"zargram/models"
)

type StoryPostgres struct {
	db *gorm.DB
}

func NewStoryPostgres(db *gorm.DB) *StoryPostgres {
	return &StoryPostgres{db: db}
}

func (p *StoryPostgres) CreateStory(story models.Story) (id int, err error) {
	story = models.Story{Title: story.Title, Content: story.Content, UserID: story.UserID}
	err = p.db.Create(&story).Error
	if err != nil {
		return 0, err
	}
	return story.ID, nil

}

func (p *StoryPostgres) GetStoriesByUserID(userID int) (stories []models.Story, err error) {
	err = p.db.Where("user_id = ?", userID).Find(&stories).Error
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (p *StoryPostgres) UpdateStoryByID(id int, t models.Story) error {
	t.ID = id
	tx := p.db.Where("id = ?", id).Save(t)

	return tx.Error
}

func (p *StoryPostgres) DeleteStoryByID(id int) (err error) {
	err = p.db.Delete(&models.Story{}, id).Error
	if err != nil {
		panic(err)
	}
	return nil
}
