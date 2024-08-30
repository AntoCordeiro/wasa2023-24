package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddToBanList(userID int, userIDToBan int) ([]types.Ban, error) {
	// Try inserting the username into the database
	_, err := db.c.Exec("INSERT INTO bansTable(userID, bannedID) VALUES (?, ?)", userID, userIDToBan)
	if err != nil {
		return nil, err
	}

	rows, err := db.c.Query("SELECT ID, userID, bannedID FROM bansTable WHERE userID = ?", userID)

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
	// Try inserting the username into the database
	_, err := db.c.Exec("DELETE FROM bansTable WHERE ID = ?", banID)
	if err != nil {
		return nil, err
	}

	rows, err := db.c.Query("SELECT ID, userID, bannedID FROM bansTable WHERE userID = ?", userID)
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
	rows, err := db.c.Query("SELECT ID, userID, bannedID FROM bansTable WHERE userID = ?", userID)
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
