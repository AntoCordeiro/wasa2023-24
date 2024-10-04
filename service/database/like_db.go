package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (db *appdbimpl) AddLike(like types.Like) (int, error) {
	// Insert the like in the likes table
	result, err := db.c.Exec("INSERT INTO likes(userID, photoID) VALUES (?, ?)", like.UserID, like.PhotoID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// Retrieve the user id and return a user object
	var likeID int
	err = db.c.QueryRow("SELECT ID FROM likes WHERE ID = ?", lastInsertID).Scan(&likeID)
	if err != nil {
		return 0, err
	}

	// Increase the likes count of the photo in the photos table
	out, err := db.c.Exec("UPDATE photos SET likesCount = likesCount + 1 WHERE ID = ?", like.PhotoID)
	if err != nil {
		return 0, err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return 0, err
	}

	return likeID, nil
}

func (db *appdbimpl) RemoveLike(userID int, photoID int) error {
	// Remove the like from the likes table
	_, err := db.c.Exec("DELETE FROM likes WHERE userID = ? AND photoID = ?", userID, photoID)
	if err != nil {
		return err
	}

	// Update the likes count of the photo in the photos table
	out, err := db.c.Exec("UPDATE photos SET likesCount = likesCount - 1 WHERE ID = ?", photoID)
	if err != nil {
		return err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return err
	}

	return nil
}