package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) StartFollowing(userID int, userIDToFollow int) ([]types.Follow, error) {
	// Try inserting the username into the database
	_, err := db.c.Exec("INSERT INTO follows(userID, followsUserID) VALUES (?, ?)", userID, userIDToFollow)
	if err != nil {
		return nil, err
	}

	rows, err := db.c.Query("SELECT ID, userID, followsUserID FROM follows WHERE userID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followsList []types.Follow
	for rows.Next() {
		var followObj types.Follow
		if err := rows.Scan(&followObj.ID, &followObj.UserID, &followObj.FollowsUserID); err != nil {
			return nil, err
		}
		followsList = append(followsList, followObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return followsList, nil
}

func (db *appdbimpl) StopFollowing(userID int, followID int) ([]types.Follow, error) {
	// Try inserting the username into the database
	_, err := db.c.Exec("DELETE FROM follows WHERE ID = ?", followID)
	if err != nil {
		return nil, err
	}

	rows, err := db.c.Query("SELECT ID, userID, followsUserID FROM follows WHERE userID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followsList []types.Follow
	for rows.Next() {
		var followObj types.Follow
		if err := rows.Scan(&followObj.ID, &followObj.UserID, &followObj.FollowsUserID); err != nil {
			return nil, err
		}
		followsList = append(followsList, followObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return followsList, nil
}

func (db *appdbimpl) GetFollowsList(userID int) ([]types.Follow, error) {
	rows, err := db.c.Query("SELECT ID, userID, followsUserID FROM follows WHERE userID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followsList []types.Follow
	for rows.Next() {
		var followObj types.Follow
		if err := rows.Scan(&followObj.ID, &followObj.UserID, &followObj.FollowsUserID); err != nil {
			return nil, err
		}
		followsList = append(followsList, followObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return followsList, nil
}
