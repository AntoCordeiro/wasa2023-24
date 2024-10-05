package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (db *appdbimpl) AddToBanList(userID int, userIDToBan int) error {
	// Try inserting the username into the database
	_, err := db.c.Exec("INSERT INTO bans(userID, bannedID) VALUES (?, ?)", userID, userIDToBan)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) RemoveFromBanList(userID int, bannedUsername string) error {
	// Delete the ban from the table
	_, err := db.c.Exec("DELETE FROM bans WHERE userID = ? AND bannedID = (SELECT ID FROM users WHERE username = ?)", userID, bannedUsername)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetBanList(userID int) ([]types.BanListComponent, error) {
	// Get and return the updated banned list
	rows, err := db.c.Query("SELECT bans.ID, username FROM bans JOIN users ON users.ID = bans.bannedID WHERE bans.userID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banList []types.BanListComponent
	for rows.Next() {
		var banObj types.BanListComponent
		if err := rows.Scan(&banObj.BanID, &banObj.Username); err != nil {
			return nil, err
		}
		banList = append(banList, banObj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return banList, nil
}
