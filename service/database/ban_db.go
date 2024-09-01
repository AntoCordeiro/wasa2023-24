package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddToBanList(userID int, userIDToBan int) ([]types.Ban, error) {
	// Try inserting the username into the database
	_, err := db.c.Exec("INSERT INTO bans(userID, bannedID) VALUES (?, ?)", userID, userIDToBan)
	if err != nil {
		return nil, err
	}

	// Stop following
	_, err = db.c.Exec("DELETE FROM follows WHERE userID = ?, followsUserID = ?", userID, userIDToBan)
	if err != nil {
		return nil, err
	}

	// Stop being followed
	_, err = db.c.Exec("DELETE FROM follows WHERE userID = ?, followsUserID = ?", userIDToBan, userID)
	if err != nil {
		return nil, err
	}

	// Delete comments of banned user under the logged in user's photos
	_, err = db.c.Exec("DELETE FROM comments WHERE userID = ? AND photoID IN (SELECT ID FROM photos WHERE userID = ?)", userIDToBan, userID)
	if err != nil {
		return nil, err
	}

	// Delete likes of banned user on the logged in user's photos
	_, err = db.c.Exec("DELETE FROM likes WHERE userID = ? AND photoID IN (SELECT ID FROM photos WHERE userID = ?)", userIDToBan, userID)
	if err != nil {
		return nil, err
	}

	// Get and return the updated banned list
	rows, err := db.c.Query("SELECT ID, userID, bannedID FROM bans WHERE userID = ?", userID)

	var banList []types.Ban
	for rows.Next() {
		var banObj types.Ban
		if err := rows.Scan(&banObj.ID, &banObj.UserID, &banObj.BannedID); err != nil {
			return nil, err
		}
		banList = append(banList, banObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return banList, nil
}

func (db *appdbimpl) RemoveFromBanList(userID int, banID int) ([]types.Ban, error) {
	// Delete the ban from the table
	_, err := db.c.Exec("DELETE FROM bans WHERE ID = ?", banID)
	if err != nil {
		return nil, err
	}

	// Get and return the updated banned list
	rows, err := db.c.Query("SELECT ID, userID, bannedID FROM bans WHERE userID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banList []types.Ban
	for rows.Next() {
		var banObj types.Ban
		if err := rows.Scan(&banObj.ID, &banObj.UserID, &banObj.BannedID); err != nil {
			return nil, err
		}
		banList = append(banList, banObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return banList, nil
}

func (db *appdbimpl) GetBanList(userID int) ([]types.Ban, error) {
	// Get and return the updated banned list
	rows, err := db.c.Query("SELECT ID, userID, bannedID FROM bans WHERE userID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banList []types.Ban
	for rows.Next() {
		var banObj types.Ban
		if err := rows.Scan(&banObj.ID, &banObj.UserID, &banObj.BannedID); err != nil {
			return nil, err
		}
		banList = append(banList, banObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return banList, nil
}
