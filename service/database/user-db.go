package database

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"

func (db *appdbimpl) GetUser(id int) (components.User, error) {
	// the function return an User(id, username) given one id

	var Id int
	var Username string

	// select the user with the given id
	err := db.c.QueryRow(`
		SELECT IdUser,Username
		FROM User
		WHERE User.IdUser=id`).Scan(&Id, &Username)

	//declare a user object
	var user components.User

	if err != nil {
		return user, nil
	}

	// if no error, assign the values retrieved
	user.Id = Id
	user.Username = Username

	// return the object
	return user, nil

}

func (db *appdbimpl) GetId(user string) (int, error) {
	// the function return the Id of the given user
	var Id int

	// select the id with the given user
	err := db.c.QueryRow(`
		SELECT IdUser
		FROM User
		WHERE user=User.Username`).Scan(&Id)

	if err != nil {
		return Id, err // to do inserire stampa di errore
	}

	return Id, nil

}
