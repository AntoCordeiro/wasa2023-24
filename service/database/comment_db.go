package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddComment(comment types.Comment) ([]types.Comment, error) {
	// Try inserting the username into the database
	_, err := db.c.Exec("INSERT INTO comments(userID, photoID, content, date) VALUES (?, ?, ?, ?)", comment.UserID, comment.PhotoID, comment.Content, comment.Date)
	if err != nil {
		return nil, err
	}

	rows, err := db.c.Query("SELECT ID, userID, photoID, content, date FROM comments WHERE photoID = ?", comment.PhotoID)
	if err != nil {
		return nil, err
	}

	var commentsList []types.Comment
	for rows.Next() {
		var commentObj types.Comment
		if err := rows.Scan(&commentObj.ID, &commentObj.UserID, &commentObj.PhotoID, &commentObj.Content, &commentObj.Date); err != nil {
			return nil, err
		}
		commentsList = append(commentsList, commentObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commentsList, nil
}

func (db *appdbimpl) RemoveComment(userID int, photoID int, commentID int) ([]types.Comment, error) {
	// Try inserting the username into the database
	_, err := db.c.Exec("DELETE FROM comments WHERE ID = ? AND userID = ?", commentID, userID)
	if err != nil {
		return nil, err
	}

	rows, err := db.c.Query("SELECT ID, userID, photoID, content, date FROM comments WHERE photoID = ?", photoID)
	if err != nil {
		return nil, err
	}

	var commentsList []types.Comment
	for rows.Next() {
		var commentObj types.Comment
		if err := rows.Scan(&commentObj.ID, &commentObj.UserID, &commentObj.PhotoID, &commentObj.Content, &commentObj.Date); err != nil {
			return nil, err
		}
		commentsList = append(commentsList, commentObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commentsList, nil
}

func (db *appdbimpl) GetCommentsList(photoID int) ([]types.Comment, error) {
	rows, err := db.c.Query("SELECT ID, userID, photoID, content, date FROM comments WHERE photoID = ?", photoID)
	if err != nil {
		return nil, err
	}

	var commentsList []types.Comment
	for rows.Next() {
		var commentObj types.Comment
		if err := rows.Scan(&commentObj.ID, &commentObj.UserID, &commentObj.PhotoID, &commentObj.Content, &commentObj.Date); err != nil {
			return nil, err
		}
		commentsList = append(commentsList, commentObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commentsList, nil
}
