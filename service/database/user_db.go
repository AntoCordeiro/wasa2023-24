package database

import (
	//"database/sql"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)



// GetName is an example that shows you how to query data
func (db *appdbimpl) UserLogin(username string) (types.User, error) {
	//try to insert the username into the databse, if it is already in the database do nothing, otherwise add it
	_, err := db.c.Exec("INSERT INTO users(username, followers, following, postCount) VALUES (?,0,0,0) ON CONFLICT (username) DO NOTHING", username)
	if err != nil {
		return types.User{}, err
	}
	//if it is already in the database take its data
	var user types.User
	if err := db.c.QueryRow("SELECT username, followers, following, postCount FROM users WHERE username = ?", username).Scan(&user.Username, &user.Followers, &user.Following, &user.PostCount); err != nil {
		return types.User{}, err
	}
	return user, nil
}


