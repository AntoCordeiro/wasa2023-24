package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err = rt.db.UserLogin(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
