package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (db *appdbimpl) UserFirstLogin(username string) (types.User, error) {
	// Insert the user in the database and get the assigned user id
	result, err := db.c.Exec("INSERT INTO users(username) VALUES (?)", username) //fix beacause if error is unique constraint it should retrieve the user anyway
	if err != nil {
		var user types.User
		if err := db.c.QueryRow("SELECT username, ID, postCount FROM users WHERE username = ?", username).Scan(&user.Username, &user.ID, &user.PostCount); err != nil {
			return user, err
		}
		return user, nil
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return types.User{}, err
	}
	// Retrieve the user id and return a user object
	var user types.User
	err = db.c.QueryRow("SELECT username, ID, postCount FROM users WHERE ID = ?", lastInsertID).Scan(&user.Username, &user.ID, &user.PostCount)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (db *appdbimpl) UserLogin(userID int, username string) (types.User, error) {
	// Try yo get the user from the database and if it exists return the user object
	var user types.User
	if err := db.c.QueryRow("SELECT ID, username, postCount FROM users WHERE ID = ? AND username = ?", userID, username).Scan(&user.ID, &user.Username, &user.PostCount); err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (db *appdbimpl) UpdateUsername(oldUsername string, newUsername string) error {
	// Update the username in the users table
	out, err := db.c.Exec("UPDATE users SET username = ? WHERE username = ?", newUsername, oldUsername)
	if err != nil {
		return err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return err
	}
	return nil
}

func (db *appdbimpl) GetProfile(profileUsername string) (types.UserProfile, error) {
	// Get the user informations from the users table
	var user types.User
	if err := db.c.QueryRow("SELECT ID, username, postCount FROM users WHERE username = ?", profileUsername).Scan(&user.ID, &user.Username, &user.PostCount); err != nil {
		return types.UserProfile{}, err
	}

	// Get the list of users that are followed by the logged in user
	followsRows, err := db.c.Query("SELECT username FROM users WHERE ID IN (SELECT followsUserID FROM follows WHERE userID = ?)", user.ID)
	if err != nil {
		return types.UserProfile{}, err
	}
	defer followsRows.Close()

	var followsList []string
	for followsRows.Next() {
		var follow string
		if err := followsRows.Scan(&follow); err != nil {
			return types.UserProfile{}, err
		}
		followsList = append(followsList, follow)
	}
	if err = followsRows.Err(); err != nil {
		return types.UserProfile{}, err
	}

	// Get the list of users who follow the logged in user
	followedRows, err := db.c.Query("SELECT username FROM users WHERE ID IN (SELECT userID FROM follows WHERE followsUserID = ?)", user.ID)
	if err != nil {
		return types.UserProfile{}, err
	}
	defer followedRows.Close()

	var followedList []string
	for followedRows.Next() {
		var followed string
		if err := followedRows.Scan(&followed); err != nil {
			return types.UserProfile{}, err
		}
		followedList = append(followedList, followed)
	}
	if err = followedRows.Err(); err != nil {
		return types.UserProfile{}, err
	}

	// Get photos uploaded by the user
	rows, err := db.c.Query("SELECT ID, userID, photoData, uploadDate, likesCount, commentsCount FROM photos WHERE userID = ? ORDER BY uploadDate DESC", user.ID)
	if err != nil {
		return types.UserProfile{}, err
	}
	defer rows.Close()

	var photosList []types.Photo
	for rows.Next() {
		var photo types.Photo
		if err := rows.Scan(&photo.ID, &photo.UserID, &photo.PhotoData, &photo.UploadDate, &photo.LikesCount, &photo.CommentsCount); err != nil {
			return types.UserProfile{}, err
		}
		var exists bool
		err = db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM likes WHERE userID = ? AND photoID = ?)", user.ID, photo.ID).Scan(&exists)
		if err != nil {
			return types.UserProfile{}, err
		}
		if exists {
			photo.IsLiked = true
		} else {
			photo.IsLiked = false
		}
		photosList = append(photosList, photo)
	}
	if err = rows.Err(); err != nil {
		return types.UserProfile{}, err
	}

	// Create and return the user profile
	return types.UserProfile{
		UserData:  user,
		Photos:    photosList,
		Follows:   followsList,
		Followers: followedList,
	}, nil
}

func (db *appdbimpl) GetID(username string) (int, error) {
	var userID int
	if err := db.c.QueryRow("SELECT ID FROM users WHERE username = ?", username).Scan(&userID); err != nil {
		return 0, err
	}
	return userID, nil
}

func (db *appdbimpl) GetStream(userID int) ([]types.Photo, error) {
	// Get the list of photos posted by users who the logged in user follows
	rows, err := db.c.Query("SELECT ID, userID, photoData, uploadDate, likesCount, commentsCount FROM photos WHERE userID IN (SELECT followsUserID FROM follows WHERE userID = ?) ORDER BY uploadDate DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MyStream []types.Photo
	for rows.Next() {
		var photo types.Photo
		if err := rows.Scan(&photo.ID, &photo.UserID, &photo.PhotoData, &photo.UploadDate, &photo.LikesCount, &photo.CommentsCount); err != nil {
			return nil, err
		}
		var exists bool
		err = db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM likes WHERE userID = ? AND photoID = ?)", userID, photo.ID).Scan(&exists)
		if err != nil {
			return nil, err
		}
		if exists {
			photo.IsLiked = true
		} else {
			photo.IsLiked = false
		}
		MyStream = append(MyStream, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return MyStream, nil
}
