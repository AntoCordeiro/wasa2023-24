package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (db *appdbimpl) AddLike(like types.Like) ([]types.Like, error) {
	// Insert the like in the likes table
	_, err := db.c.Exec("INSERT INTO likes(userID, photoID, date) VALUES (?, ?, ?)", like.UserID, like.PhotoID, like.Date)
	if err != nil {
		return nil, err
	}

	// Increase the likes count of the photo in the photos table
	out, err := db.c.Exec("UPDATE photos SET likesCount = likesCount + 1 WHERE ID = ?", like.PhotoID)
	if err != nil {
		return nil, err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return nil, err
	}

	// Get and return the updated likes list
	rows, err := db.c.Query("SELECT ID, userID, photoID, date FROM likes WHERE userID = ? AND photoID = ? ORDER BY date DESC", like.UserID, like.PhotoID)
	if err != nil {
		return nil, err
	}

	var likesList []types.Like
	for rows.Next() {
		var likeObj types.Like
		if err := rows.Scan(&likeObj.ID, &likeObj.UserID, &likeObj.PhotoID, &likeObj.Date); err != nil {
			return nil, err
		}
		likesList = append(likesList, likeObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return likesList, nil
}

func (db *appdbimpl) RemoveLike(likeID int, userID int, photoID int) ([]types.Like, error) {
	// Remove the like from the likes table
	_, err := db.c.Exec("DELETE FROM likes WHERE ID = ?", likeID)
	if err != nil {
		return nil, err
	}

	// Update the likes count of the photo in the photos table
	out, err := db.c.Exec("UPDATE photos SET likesCount = likesCount - 1 WHERE ID = ?", photoID)
	if err != nil {
		return nil, err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return nil, err
	}

	// Get and return the updated likes list
	rows, err := db.c.Query("SELECT ID, userID, photoID, date FROM likes WHERE userID = ? AND photoID = ? ORDER BY date DESC", userID, photoID)
	if err != nil {
		return nil, err
	}

	var likesList []types.Like
	for rows.Next() {
		var likeObj types.Like
		if err := rows.Scan(&likeObj.ID, &likeObj.UserID, &likeObj.PhotoID, &likeObj.Date); err != nil {
			return nil, err
		}
		likesList = append(likesList, likeObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return likesList, nil
}

func (db *appdbimpl) GetLikesList(userID int, photoID int) ([]types.Like, error) {
	// Get and return the updated likes list
	rows, err := db.c.Query("SELECT ID, userID, photoID, date FROM likes WHERE userID = ? AND photoID = ? ORDER BY date DESC", userID, photoID)
	if err != nil {
		return nil, err
	}

	var likesList []types.Like
	for rows.Next() {
		var likeObj types.Like
		if err := rows.Scan(&likeObj.ID, &likeObj.UserID, &likeObj.PhotoID, &likeObj.Date); err != nil {
			return nil, err
		}
		likesList = append(likesList, likeObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return likesList, nil
}
