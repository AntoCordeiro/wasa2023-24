/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// user operations
	UserFirstLogin(username string) (types.User, error)
	UserLogin(userID int, username string) (types.User, error)
	UpdateUsername(oldUsername string, newUsername string) error
	GetProfile(profileUsername string) (types.UserProfile, error)
	GetID(username string) (int, error)
	GetStream(userID int) ([]types.Photo, error)

	// photo operations
	InsertPhoto(photoObj types.Photo) error
	RemovePhoto(userID int, photoID int) error

	// follow operations
	StartFollowing(userID int, userIDToFollow int) error
	StopFollowing(userID int, followedUsername string) error

	// ban operations
	GetBanList(userID int) ([]types.BanListComponent, error)
	AddToBanList(userID int, userIDToBan int) error
	RemoveFromBanList(userID int, bannedUsername string) error

	// likes operations
	GetLikesList(userID int, photoID int) ([]types.LikeListComponent, error)
	AddLike(like types.Like) (int, error)
	RemoveLike(likeID int, userID int, photoID int) error

	// comments operations
	GetCommentsList(photoID int) ([]types.CommentListComponent, error)
	AddComment(comment types.Comment) (int, error)
	RemoveComment(userID int, photoID int, commentID int) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		usersTable := `CREATE TABLE users (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						username VARCHAR(16) UNIQUE CHECK (LENGTH(username) BETWEEN 3 AND 16),
						postCount INTEGER DEFAULT 0);`
		_, err = db.Exec(usersTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Create the photos table
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		photosTable := `CREATE TABLE photos (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						userID  VARCHAR(16),
						photoData BLOB,
						uploadDate DATETIME,
						likesCount INTEGER,
						commentsCount INTEGER,
						FOREIGN KEY (userID) REFERENCES users(ID));`
		_, err = db.Exec(photosTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Create the follows table
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='followsTable';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		followsTable := `CREATE TABLE follows (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						userID  INTEGER,
						followsUserID INTEGER,
						FOREIGN KEY (userID) REFERENCES users(ID),
						FOREIGN KEY (followsUserID) REFERENCES users(ID));`
		_, err = db.Exec(followsTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Create the bans table
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='bansTable';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		bansTable := `CREATE TABLE bans (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						userID  INTEGER,
						bannedID INTEGER,
						FOREIGN KEY (userID) REFERENCES users(ID),
						FOREIGN KEY (bannedID) REFERENCES users(ID));`
		_, err = db.Exec(bansTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Create the likes table
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		likesTable := `CREATE TABLE likes (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						userID  INTEGER,
						photoID INTEGER,
						date DATETIME,
						FOREIGN KEY (userID) REFERENCES users(ID),
						FOREIGN KEY (photoID) REFERENCES photos(ID));`
		_, err = db.Exec(likesTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Create the comments table
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		commentsTable := `CREATE TABLE comments (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						userID  INTEGER,
						photoID INTEGER,
						content TEXT,
						date DATETIME,
						FOREIGN KEY (userID) REFERENCES users(ID),
						FOREIGN KEY (photoID) REFERENCES photos(ID));`
		_, err = db.Exec(commentsTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
