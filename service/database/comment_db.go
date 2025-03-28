package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (db *appdbimpl) AddComment(comment types.Comment) (int, error) {
	// Insert the new comment in the comments table
	result, err := db.c.Exec("INSERT INTO comments(userID, photoID, content, date) VALUES (?, ?, ?, ?)", comment.UserID, comment.PhotoID, comment.Content, comment.Date)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// Retrieve the user id and return a user object
	var commentID int
	err = db.c.QueryRow("SELECT ID FROM comments WHERE ID = ?", lastInsertID).Scan(&commentID)
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

func (db *appdbimpl) RemoveComment(userID int, photoID int, commentID int) error {
	// Try inserting the username into the database
	_, err := db.c.Exec("DELETE FROM comments WHERE ID = ? AND userID = ?", commentID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetCommentsList(myUserID int, photoID int) ([]types.CommentListComponent, error) {
	rows, err := db.c.Query("SELECT comments.ID, userID, photoID, content, date, username FROM comments JOIN users ON users.ID = comments.userID WHERE comments.photoID = ? AND userID NOT IN (SELECT bannedID FROM bans WHERE userID = ?)", photoID, myUserID)
	if err != nil {
		return nil, err
	}

	var commentsList []types.CommentListComponent
	for rows.Next() {
		var commentObj types.CommentListComponent
		if err := rows.Scan(&commentObj.CommentData.ID, &commentObj.CommentData.UserID, &commentObj.CommentData.PhotoID, &commentObj.CommentData.Content, &commentObj.CommentData.Date, &commentObj.CommentData.Username); err != nil {
			return nil, err
		}
		commentsList = append(commentsList, commentObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commentsList, nil
}
