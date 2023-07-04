package repository

import (
	"gorm.io/gorm"
	"zargram/models"
)

type PostPostgres struct {
	db *gorm.DB
}

func NewPostPostgres(db *gorm.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (p *PostPostgres) CreatePost(post models.Post) (id int, err error) {
	post = models.Post{Title: post.Title, Content: post.Content, UserID: post.UserID}
	err = p.db.Create(&post).Error
	if err != nil {
		return 0, err
	}
	return post.ID, nil

}

func (p *PostPostgres) GetPostsByUserID(userID int) (posts []models.Post, err error) {
	err = p.db.Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostPostgres) UpdatePostByID(id int, t models.Post) error {
	t.ID = id
	tx := p.db.Where("id = ?", id).Save(t)

	return tx.Error
}

func (p *PostPostgres) DeletePostByID(id int) (err error) {
	err = p.db.Delete(&models.Post{}, id).Error
	if err != nil {
		panic(err)
	}
	return nil
}

func (p *PostPostgres) MarkArchived(id int, t models.Post) error {
	t.Archived = true
	if err := p.db.Save(&t).Error; err != nil {
		return err
	}
	return nil
}
func (p *PostPostgres) GetPostByID(id int) (t models.Post, err error) {
	if err = p.db.Where("id=?", id).Take(&t).Error; err != nil {
		return models.Post{}, err
	}
	return t, nil
}
