package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)



// GetName is an example that shows you how to query data
func (db *appdbimpl) StartFollowing(username string, usernameToFollow string) ([]types.Follow, error) {
    // Try inserting the username into the database
    _, err := db.c.Exec("INSERT INTO followsTable(username, followsUsername) VALUES (?, ?)", username, usernameToFollow)
    if err != nil {
		return nil, err
	}

    rows, err := db.c.Query("SELECT ID, username, followsUsername FROM followsTable WHERE username = ?", username)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var followsList []types.Follow
    for rows.Next() {
        var followObj types.Follow
        if err := rows.Scan(&followObj.ID, &followObj.Username, &followObj.FollowsUsername); err != nil {
            return nil, err
        }
        followsList = append(followsList, followObj)
    }

    return followsList, nil
}

func (db *appdbimpl) StopFollowing(username string, followID int) ([]types.Follow, error) {
    // Try inserting the username into the database
    _, err := db.c.Exec("DELETE FROM followsTable WHERE ID = ?", followID)
    if err != nil {
        return nil, err
    }

    rows, err := db.c.Query("SELECT ID, username, followsUsername FROM followsTable WHERE username = ?", username)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var followsList []types.Follow
    for rows.Next() {
        var followObj types.Follow
        if err := rows.Scan(&followObj.ID, &followObj.Username, &followObj.FollowsUsername); err != nil {
            return nil, err
        }
        followsList = append(followsList, followObj)
    }

    return followsList, nil
}

func (db *appdbimpl) GetFollowsList(username string) ([]types.Follow, error) {
    rows, err := db.c.Query("SELECT ID, username, followsUsername FROM followsTable WHERE username = ?", username)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var followsList []types.Follow
    for rows.Next() {
        var followObj types.Follow
        if err := rows.Scan(&followObj.ID, &followObj.Username, &followObj.FollowsUsername); err != nil {
            return nil, err
        }
        followsList = append(followsList, followObj)
    }

    return followsList, nil
}