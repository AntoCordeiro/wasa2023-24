package types

import ("time")

type User struct {
    Username    string `json:"username"`
    ID          int    `json:"id"`
    Followers   int    `json:"followers"`
    Following   int    `json:"following"`
    PostCount   int    `json:"postCount"`
}

type Photo struct {
	ID              int        		`json:"id"`
	Username        string     		`json:"username"`
	PhotoData       []byte     		`json:"photoData"`
	UploadDate      time.Time		`json:"uploadDate"`
	LikesCount      int        		`json:"likesCount"`
	CommentsCount   int        		`json:"commentsCount"`
}

type UserProfile struct {
    UserData  User
    Photos []Photo
}

type Follow struct {
    ID                  int         `json:"id"`
    Username            string      `json:"username"`
    FollowsUsername     string      `json:"followsUsername"`
}