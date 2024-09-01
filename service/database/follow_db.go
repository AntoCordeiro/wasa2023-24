package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (db *appdbimpl) StartFollowing(userID int, userIDToFollow int) ([]types.Follow, error) {
	// Insert the new follow in the follows table
	_, err := db.c.Exec("INSERT INTO follows(userID, followsUserID) VALUES (?, ?)", userID, userIDToFollow)
	if err != nil {
		return nil, err
	}

	// Get and return the updated list of follows
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
	// Delete the follow from the follows table
	_, err := db.c.Exec("DELETE FROM follows WHERE ID = ?", followID)
	if err != nil {
		return nil, err
	}

	// Get and return the updated follows list
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
	// Get and return the updated follows list
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
