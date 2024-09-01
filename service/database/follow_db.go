package database

func (db *appdbimpl) StartFollowing(userID int, userIDToFollow int) error {
	// Insert the new follow in the follows table
	_, err := db.c.Exec("INSERT INTO follows(userID, followsUserID) VALUES (?, ?)", userID, userIDToFollow)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) StopFollowing(userID int, followedUsername string) error {
	// Delete the follow from the follows table
	out, err := db.c.Exec("DELETE FROM follows WHERE followsUserID = (SELECT ID FROM users WHERE username = ?) AND userID", followedUsername, userID)
	if err != nil {
		return err
	}
	affectedRows, err := out.RowsAffected()
	if err != nil || affectedRows == 0 {
		return err
	}
	return nil
}
