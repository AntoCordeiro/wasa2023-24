package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (db *appdbimpl) InsertPhoto(photoObj types.Photo) error {
	// Insert the photo in the photos table
	_, err := db.c.Exec("INSERT INTO photos(userID, photoData, uploadDate, likesCount, commentsCount) VALUES (?, ?, ?, ?, ?)", photoObj.UserID, photoObj.PhotoData, photoObj.UploadDate, photoObj.LikesCount, photoObj.CommentsCount) //fix beacause if error is unique constraint it should retrieve the user anyway

	// Update the posts count of the users in the users table
	out, err := db.c.Exec("UPDATE users SET postCount = postCount + 1 WHERE ID = ?", photoObj.UserID)
	if err != nil {
		return err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return err
	}

	return nil
}

func (db *appdbimpl) RemovePhoto(userID int, photoID int) error {
	// Remove the photo from the photos table
	_, err := db.c.Exec("DELETE FROM photos WHERE userID = ? AND ID = ?", userID, photoID)
	if err != nil {
		return err
	}

	// Update the posts count of the users in the users table
	out, err := db.c.Exec("UPDATE users SET postCount = postCount - 1 WHERE ID = ?", userID)
	if err != nil {
		return err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return err
	}

	// Delete the comments under the deleted photo 
	_, err = db.c.Exec("DELETE FROM comments WHERE photoID = ?", photoID)
	if err != nil {
		return err
	}

	// Delete the likes on the the deleted photo
	_, err = db.c.Exec("DELETE FROM likes WHERE photoID", photoID)
	if err != nil {
		return err
	}

	return nil
}
