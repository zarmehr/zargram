package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name" gorm:"not null" `
	Username string `json:"username" gorm:"unique;size:32;not null"`
	Password string `json:"password_hash" gorm:"not null"`
	Avatar   string `json:"-"`
	//Posts    []Post   `json:"-"`
	//Friends  []Friend `gorm:"foreignKey:UserID" json:"friends"`
}

func (u User) HasAvatar() bool {
	return u.Avatar != ""
}

// Post represents the post table in the database
type Post struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"size:255" json:"title"`
	Content  string `gorm:"type:text" json:"content"`
	Archived bool   `gorm:"default:false" json:"archived"`
	UserID   int    `gorm:"index" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"-"`
	//Comments  []Comment  `json:"comments"`
	//Reactions []Reaction `gorm:"foreignKey:PostID"json:"reactions"`
}

// Comment represents the comment table in the database
type Comment struct {
	ID              int    `gorm:"primaryKey" json:"id"`
	PostID          int    `gorm:"index" json:"post_id"`
	ParentCommentID int    `gorm:"index" json:"parent_comment_id"`
	Content         string `gorm:"type:text" json:"content"`
	//Replies         []Comment `gorm:"foreignKey:ParentCommentID" json:"replies"`
	UserID int  `gorm:"index" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
}

// Story represents the story table in the database
type Story struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"size:255" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `gorm:"index" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"-"`
}

// Reaction represents the reaction table in the database
type Reaction struct {
	ID     int `gorm:"primaryKey" json:"id"`
	PostID int `gorm:"index" json:"post_id"`
	//CommentID int    `gorm:"index" json:"comment_id"`
	UserID int    `gorm:"index" json:"user_id"`
	Type   string `json:"type"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
}
type Friend struct {
	ID       int  `gorm:"primaryKey" json:"id"`
	UserID   int  `json:"userId"`
	User     User `json:"-"`
	FriendID int  `json:"friendId"`
	Friend   User `json:"-"`
}
