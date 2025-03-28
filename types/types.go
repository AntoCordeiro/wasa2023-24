package types

import (
	"time"
)

type User struct {
	Username  string `json:"username"`
	ID        int    `json:"id"`
	PostCount int    `json:"postCount"`
}

type Photo struct {
	ID            int       `json:"id"`
	UserID        int       `json:"userID"`
	Username      string    `json:"username"`
	PhotoData     []byte    `json:"photoData"`
	UploadDate    time.Time `json:"uploadDate"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	IsLiked       bool      `json:"isLiked"`
}

type UserProfile struct {
	UserData  User     `json:"user"`
	Photos    []Photo  `json:"photos"`
	Follows   []string `json:"follows"`
	Followers []string `json:"followers"`
}

type Follow struct {
	ID            int `json:"id"`
	UserID        int `json:"userID"`
	FollowsUserID int `json:"followsUserID"`
}

type Ban struct {
	ID       int `json:"id"`
	UserID   int `json:"userID"`
	BannedID int `json:"bannedID"`
}

type BanListComponent struct {
	BanID    int    `json:"banID"`
	Username string `json:"username"`
}

type Like struct {
	ID      int `json:"id"`
	UserID  int `json:"userID"`
	PhotoID int `json:"photoID"`
}

type LikeListComponent struct {
	LikeID   int    `json:"likeID"`
	Username string `json:"username"`
}

type Comment struct {
	ID       int       `json:"id"`
	UserID   int       `json:"userID"`
	Username string    `json:"username"`
	PhotoID  int       `json:"photoID"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
}

type CommentListComponent struct {
	CommentData Comment `json:"commentData"`
	Username    string  `json:"username"`
}
