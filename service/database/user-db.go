package database
import  "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"

func (db *appdbimpl) GetUser(id int) (components.User, error) {

	var Id int
	var Username string

	// select the user with the given id
	err := db.c.QueryRow(`
		SELECT Id,Username
		FROM User
		WHERE Id.User=id`).Scan(&Id, &Username)

	//declare a user object
	var user components.User

	if err != nil {
		return user, nil
	}

	// if no error, assign the values retrievede
	user.Id = Id
	user.Username = Username

	// return the object
	return user, nil

}