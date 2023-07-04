package repository

import (
	"gorm.io/gorm"
	"zargram/models"
)

type ReactionPostgres struct {
	db *gorm.DB
}

func NewReactionPostgres(db *gorm.DB) *ReactionPostgres {
	return &ReactionPostgres{db: db}
}

func (p *ReactionPostgres) CreateReaction(reaction models.Reaction) (id int, err error) {
	reaction = models.Reaction{PostID: reaction.PostID, UserID: reaction.UserID, Type: reaction.Type}
	err = p.db.Create(&reaction).Error
	if err != nil {
		return 0, err
	}
	return reaction.ID, nil

}
func (p *ReactionPostgres) GetReactionsByPostID(postID int) (reactions []models.Reaction, err error) {
	err = p.db.Where("post_id = ?", postID).Find(&reactions).Error
	if err != nil {
		return nil, err
	}
	return reactions, nil
}

func (p *ReactionPostgres) UpdateReactionByID(id int, t models.Reaction) error {
	t.ID = id
	tx := p.db.Where("id = ?", id).Save(t)

	return tx.Error
}

func (p *ReactionPostgres) DeleteReactionByID(id int) (err error) {
	err = p.db.Delete(&models.Reaction{}, id).Error
	if err != nil {
		panic(err)
	}
	return nil
}
