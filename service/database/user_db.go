package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)



// GetName is an example that shows you how to query data
func (db *appdbimpl) UserFirstLogin(username string) (types.User, error) {
    // Try inserting the username into the database
    result, err := db.c.Exec("INSERT INTO users(username) VALUES (?)", username)  //fix beacause if error is unique constraint it should retrieve the user anyway
    if err != nil {
		var user types.User
		if err := db.c.QueryRow("SELECT username, ID, followers, following, postCount FROM users WHERE username = ?", username).Scan(&user.Username, &user.ID, &user.Followers, &user.Following, &user.PostCount); err != nil {
			return user, err
		}
		return user, nil
	}
    // Get the last insert ID
    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return types.User{}, err
    }
    // Retrieve the inserted data using the last insert ID
    var user types.User
    err = db.c.QueryRow("SELECT username, ID, followers, following, postCount FROM users WHERE ID = ?", lastInsertID).Scan(&user.Username, &user.ID, &user.Followers, &user.Following, &user.PostCount)
    if err != nil {
        return types.User{}, err
    }
    return user, nil
}

func (db *appdbimpl) UserLogin(userID int) (types.User, error) {
	//try yo get the user from the database and if it exists return the user object
	var user types.User
	if err := db.c.QueryRow("SELECT ID, username, followers, following, postCount FROM users WHERE ID = ?", userID).Scan(&user.ID, &user.Username, &user.Followers, &user.Following, &user.PostCount); err != nil {
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
	if err := db.c.QueryRow("SELECT username, followers, following, postCount FROM users WHERE username = ?", profileUsername).Scan(&user.Username, &user.Followers, &user.Following, &user.PostCount); err != nil {
		return types.UserProfile{}, err
	}
	// Get photos uploaded by the user
    var photos []types.Photo
    rows, err := db.c.Query("SELECT ID, username, photoData, uploadDate, likesCount, commentsCount FROM photos WHERE username = ?", user.Username)
    if err != nil {
        return types.UserProfile{}, err
    }
    defer rows.Close()
	// Add each photo to the slice   
	for rows.Next() {
        var photo types.Photo 
        if err := rows.Scan(&photo.ID, &photo.Username, &photo.PhotoData, &photo.UploadDate, &photo.LikesCount, &photo.CommentsCount); err != nil {
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

