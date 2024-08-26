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
	UserLogin(username string) (types.User, error)
	UpdateUsername(newUsername string) (error)
	GetProfile(profileUsername string) (types.UserProfile, error)	
	
	GetName() (string, error)
	SetName(name string) error

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
						followers INTEGER DEFAULT 0,
						following INTEGER DEFAULT 0,
						postCount INTEGER DEFAULT 0);`
		_, err = db.Exec(usersTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		photosTable := `CREATE TABLE photos (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						userID INTEGER NOT NULL,
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

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
