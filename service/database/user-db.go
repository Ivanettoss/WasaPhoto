package database

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
)

func (db *appdbimpl) GetUser(id int) (components.User, error) {
	// the function return an User(id, username) given one id

	var Id int
	var Username string

	// select the user with the given id
	err := db.c.QueryRow(`
		SELECT IdUser,Username
		FROM User
		WHERE User.IdUser=?`, id).Scan(&Id, &Username)

	// declare a user object
	var user components.User

	if err != nil {
		return user, err
	}

	// if no error, assign the values retrieved
	user.Id = Id
	user.Username = Username

	// return the object
	return user, err

}

func (db *appdbimpl) GetId(user string) (int, error) {
	// the function return the Id of the given user
	var Id int

	// select the id with the given user
	err := db.c.QueryRow(`
		SELECT IdUser
		FROM User
		WHERE User.Username=?`, user).Scan(&Id)

	if err != nil {
		return Id, err // to do inserire stampa di errore
	}

	return Id, nil

}

func (db *appdbimpl) SetUsername(username string, new_username string) (err error) {
	_, err = db.c.Exec(`UPDATE User SET Username = ? WHERE  Username = ?`, new_username, username)

	if err != nil {
		return err
	}

	return err

}

func (db *appdbimpl) SearchUsername(myName string, username string) (components.Ulist, error) {
	users := components.Ulist{}
	users.UsersList = []components.Username{}
	myId, err := db.GetId(myName)
	if err != nil {
		return users, ErrUserNotFound
	}
	userRows, err := db.c.Query(`
		SELECT Username
		FROM User
		WHERE User.Username LIKE '%'||?||'%' and User.Username!=?
		EXCEPT 
			SELECT User.Username
			FROM Ban,User
			WHERE Ban.IdUser=User.IdUser and IdUserBanned=?`, username, myName, myId)

	if err != nil {
		return users, err
	}

	defer userRows.Close()

	for userRows.Next() {
		var name components.Username
		err = userRows.Scan(&name.Username)

		if err != nil {

			return users, err
		}

		users.UsersList = append(users.UsersList, name)
	}
	if userRows.Err() != nil {
		return users, err
	}

	// format the list
	return users, err

}
