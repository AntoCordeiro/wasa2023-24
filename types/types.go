package types

type Date struct {
	Year  int        // Year (e.g., 2014).
	Month int // Month of the year (January = 1, ...).
	Day   int        // Day of the month, starting at 1.
}

type Time struct {
	Hour       int // The hour of the day in 24-hour format; range [0-23]
	Minute     int // The minute of the hour; range [0-59]
	Second     int // The second of the minute; range [0-59]
}

type DateTime struct {
	Date Date
	Time Time
}

type User struct {
    Username    string `json:"username"`
    ID          int    `json:"id"`
    Followers   int    `json:"followers"`
    Following   int    `json:"following"`
    PostCount   int    `json:"postCount"`
}

type Photo struct {
	ID              int        `json:"id"`
	UserID          uint64     `json:"userId"`
	PhotoData       []byte     `json:"photoData"`
	UploadDate      DateTime   `json:"uploadDate"`
	LikesCount      int        `json:"likesCount"`
	CommentsCount   int        `json:"commentsCount"`
}

type UserProfile struct {
    UserData  User
    Photos []Photo
}