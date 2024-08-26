package database

import (
	//"database/sql"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
	"fmt"
)



// GetName is an example that shows you how to query data
func (db *appdbimpl) UserFirstLogin(username string) (types.User, error) {
    // Try inserting the username into the database
    result, err := db.c.Exec("INSERT INTO users(username) VALUES (?) ON CONFLICT (username) DO NOTHING", username)
    if err != nil {
        return types.User{}, err
    }

    // Check if any rows were affected by the insert
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return types.User{}, err
    }

    // If no rows were affected, a user with the same username might already exist
    if rowsAffected == 0 {
        fmt.Println("Warning: Username", username, "already exists, not inserting new user.")
        // You can choose to return an error or handle the existing user case differently
    }

    // Query to get the user data
    var user types.User
    if err := db.c.QueryRow("SELECT username, ID, followers, following, postCount FROM users WHERE username = ?", username).Scan(&user.Username, &user.ID, &user.Followers, &user.Following, &user.PostCount); err != nil {
        return types.User{}, err
    }

    return user, nil
}

func (db *appdbimpl) UserLogin(username string) (types.User, error) {
	//try yo get the user from the database and if it exists return the user object
	var user types.User
	if err := db.c.QueryRow("SELECT ID, username, followers, following, postCount FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Followers, &user.Following, &user.PostCount); err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (db *appdbimpl) UpdateUsername(newUsername string) (error) {
	out, err := db.c.Exec("UPDATE users SET username = ? WHERE username = ?", newUsername)
	if err != nil {
		return err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0{
		return err
	}
	return nil
}

func (db *appdbimpl) GetProfile(profileUsername string) (types.UserProfile, error) {
	var user types.User
	if err := db.c.QueryRow("SELECT username, ID, followers, following, postCount FROM users WHERE username = ?", profileUsername).Scan(&user.Username, &user.ID, &user.Followers, &user.Following, &user.PostCount); err != nil {
		return types.UserProfile{}, err
	}
	// Get photos uploaded by the user
    var photos []types.Photo
    rows, err := db.c.Query("SELECT ID, userID, photoData, uploadDate, likesCount, commentsCount FROM photos WHERE user_id = ?", user.ID)
    if err != nil {
        return types.UserProfile{}, err
    }
    defer rows.Close()
	// Add each photo to the slice   
	for rows.Next() {
        var photo types.Photo 
        if err := rows.Scan(&photo.ID, &photo.UserID, &photo.PhotoData, &photo.UploadDate, &photo.LikesCount, &photo.CommentsCount); err != nil {
            return types.UserProfile{}, err
        }
        photos = append(photos, photo)
    }

    // Create and return the user profile
    return types.UserProfile{
        UserData:  user,
        Photos: photos,
    }, nil
}

