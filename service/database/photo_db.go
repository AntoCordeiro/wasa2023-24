package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)



// GetName is an example that shows you how to query data
func (db *appdbimpl) InsertPhoto(photoObj types.Photo) (error) {
    // Try inserting the username into the database
    _, err := db.c.Exec("INSERT INTO photos(username, photoData, uploadDate, likesCount, commentsCount) VALUES (?, ?, ?, ?, ?)", photoObj.Username, photoObj.PhotoData, photoObj.UploadDate, photoObj.LikesCount, photoObj.CommentsCount)  //fix beacause if error is unique constraint it should retrieve the user anyway
    return err
}
