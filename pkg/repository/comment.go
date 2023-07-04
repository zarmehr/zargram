package repository

import (
	"gorm.io/gorm"
	"zargram/models"
)

type CommentPostgres struct {
	db *gorm.DB
}

func NewCommentPostgres(db *gorm.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (c *CommentPostgres) CreateComment(comment models.Comment) (int, error) {
	comment = models.Comment{PostID: comment.PostID, ParentCommentID: comment.ParentCommentID, Content: comment.Content, UserID: comment.UserID}

	err := c.db.Create(&comment).Error
	if err != nil {
		return 0, err
	}
	return comment.ID, nil

}
func (c *CommentPostgres) GetCommentsByPostID(postID int) (comments []models.Comment, err error) {
	err = c.db.Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (c *CommentPostgres) UpdateCommentByID(id int, t models.Comment) error {
	t.ID = id
	tx := c.db.Where("id = ?", id).Save(t)

	return tx.Error

}

func (c *CommentPostgres) DeleteCommentByID(id int) (err error) {
	var comment models.Comment
	//find comment we want to delete
	err = c.db.Where("id=?", id).Delete(&comment).Error
	if err != nil {
		return err
	}

	//Now let's see if it had replies
	var comment2 models.Comment
	err = c.db.Where("parent_comment_id=?", id).Find(&comment2).Error
	if err != nil {
		return err
	}
	//Recursively deletes replies

	if comment2.ID != 0 {
		err2 := c.DeleteCommentByID(comment2.ID)
		if err2 != nil {
			return err2
		}
	}
	return nil
}
