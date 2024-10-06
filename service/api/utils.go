package api

import (
	"net/http"
	"regexp"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
	"github.com/julienschmidt/httprouter"
)

func GetBearerToken(authRaw string) (int, error) {

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	tokenString := re.FindAllString(authRaw, -1)

	if len(tokenString) == 0 {
		return -1, nil
	}
	token, err := strconv.Atoi(tokenString[0])

	return token, err
}

func (rt *_router) UserAuthentication(username string, w http.ResponseWriter, r *http.Request, ps httprouter.Params) (components.User, error) {
	// the function authenticate the user given  the username aggiungi il codice errore http
	var user components.User
	var idFromDB int

	// autenticazione
	uname := ps.ByName(username) // return string (name in risorsa)

	// query per ottenere id di utente user
	idFromDB, err := rt.db.GetId(uname)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return user, err
	}

	authString := r.Header.Get("Authorization") // prende nella req il pezzo col barer

	// estraggo il token id dal barer
	token, err := GetBearerToken(authString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return user, err
	}

	if idFromDB == token {
		// coincidono, posso autenticare e returnare l'utente

		user.Id = idFromDB
		user.Username = uname
		return user, err

	} else {
		return user, err // to do return the errore corretto
	}

}
func CheckAuth(user components.User, authRaw string) error {
	token, err := GetBearerToken(authRaw)

	if err != nil {
		return err
	}

	if int(user.Id) != token {
		return err
	}

	return nil
}
