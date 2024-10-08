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

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	InsertUser(name string) (int, error)
	GetUser(id int) (components.User, error)
	GetId(user string) (int, error)

	InsertFollow(idUserPerforming int, idUserToFollow int) error
	DeleteFollow(idUserPerforming int, idUserToUnFollow int) error

	InsertBan(idUserPerforming int, idUserToBan int) error
	DeleteBan(idUserPerforming int, idUserToUnBan int) error
	BanCheck(OwnerId int, userPerforming int) (bool, error)

	InsertLike(idUserLike int, idPhoto int) error
	GetLike(idUserPerforming int, idPhoto int) error
	DeleteLike(idUserPerforming int, idPhoto int) error
	CountLikes(idPhoto int) (int, error)
	CheckLike(idUserPerforming int, idPhoto int) (bool, error)

	InsertComment(comment components.Comment) (int, error)
	GetComment(commentId int) (components.Comment, error)
	DeleteComment(commentId int) error

	GetCommentList(idPhoto int) ([]components.Comment, error)

	GetFollowedList(idUser int) ([]string, error)
	GetFollowersList(idUser int) ([]string, error)
	GetFollowedNumber(idUser int) (int, error)
	GetFollowersNumber(idUser int) (int, error)
	GetFollowState(userPerformingId int, profileOwnerId int) (bool, error)

	GetBannedList(idUserPerforming int) ([]string, error)

	InsertPhoto(IdUser int, photo components.Photo) error
	DeletePhoto(photoId int) error
	GetPhoto(photoId int) (components.Photo, error)
	GetPhotoOwnerId(photoId int) (int, error)
	GetPostedPhotoNumber(idUser int) (int, error)
	GetUserPhotos(idUserPerforming int, username string) ([]components.Photo, error)
	GetStream(userPerformingId int, userPerformingName string) (components.Stream, error)
	SetUsername(username string, new_username string) (err error)
	SearchUsername(myName string, username string) (components.Ulist, error)
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

	var err error
	_, err = db.Exec("PRAGMA foreign_key=ON")
	if err != nil {
		return nil, err
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	User := `
		CREATE TABLE IF NOT EXISTS User (
			IdUser INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE 
		);
		`

	Photo := ` 
			CREATE TABLE IF NOT EXISTS Photo (
				IdPhoto INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				IdUser INTEGER NOT NULL,  
			    DateTime TEXT NOT NULL,
				Path TEXT NOT NULL,
				FOREIGN KEY (IdUser) REFERENCES User(Username)
			)
			`

	Like := ` 
		CREATE TABLE  IF NOT EXISTS Like (
			IdUser INTEGER NOT NULL, 
			IdPhoto INTEGER NOT NULL,  
			PRIMARY KEY (IdPhoto, IdUser),
			FOREIGN KEY	 (IdPhoto) REFERENCES Photo,
			FOREIGN KEY  (IdUser) REFERENCES  User 
		
	)
	`
	Follow := ` 
		CREATE TABLE  IF NOT EXISTS Follow(
			IdUser INTEGER NOT NULL ,
			IdUserFollowed INTEGER NOT NULL,  
			PRIMARY KEY (IdUser, IdUserFollowed),
			FOREIGN KEY	 (IdUser) REFERENCES User,
			FOREIGN KEY  (IdUserFollowed) REFERENCES  User 
		
	)
	`
	Comment := ` 
		CREATE TABLE  IF NOT EXISTS Comment(
			IdComment INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			IdUser INTEGER NOT NULL,
			IdPhoto INTEGER NOT NULL,  
			DataTime TEXT NOT NULL ,
			Text TEXT NOT NULL,
			FOREIGN KEY	 (IdUser) REFERENCES User,
			FOREIGN KEY  (IdPhoto) REFERENCES  User 
		
	)
	`

	Ban := ` 
		CREATE TABLE  IF NOT EXISTS Ban(
			
			IdUser INTEGER NOT NULL, 
			IdUserBanned INTEGER NOT NULL,  
			PRIMARY KEY (IdUser, IdUserBanned),
			FOREIGN KEY	 (IdUser) REFERENCES User,
			FOREIGN KEY  (IdUserBanned) REFERENCES  User 
		
	)
	`

	Upload := ` 
		CREATE TABLE  IF NOT EXISTS Upload(
			IdUser INTEGER NOT NULL, 
			IdPhoto INTEGER NOT NULL,  
			DataTime TEXT NOT NULL,
			FOREIGN KEY	 (IdUser) REFERENCES User,
			FOREIGN KEY  (IdPhoto) REFERENCES  User 
		
	)
	`

	_, err = db.Exec(User)

	if err != nil {
		return nil, fmt.Errorf("Error in User table creation: %w", err)
	}

	_, err = db.Exec(Photo)
	if err != nil {
		return nil, fmt.Errorf("Error in  Photo table creation: %w", err)
	}

	_, err = db.Exec(Like)
	if err != nil {
		return nil, fmt.Errorf("Error in Like table creation: %w", err)
	}

	_, err = db.Exec(Follow)
	if err != nil {
		return nil, fmt.Errorf("Error in Follow table creation: %w", err)
	}

	_, err = db.Exec(Comment)
	if err != nil {
		return nil, fmt.Errorf("Error in Comment table creation: %w", err)
	}

	_, err = db.Exec(Ban)
	if err != nil {
		return nil, fmt.Errorf("Error in Ban table creation: %w", err)
	}

	_, err = db.Exec(Upload)
	if err != nil {
		return nil, fmt.Errorf("Error in Upload table creation: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil

}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
