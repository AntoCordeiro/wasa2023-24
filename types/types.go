package types

type User struct {
    Username string `json:"username"`
    Followers int   `json:"followers"`
    Following int   `json:"following"`
    PostCount int   `json:"postCount"`
}