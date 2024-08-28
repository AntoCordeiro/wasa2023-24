package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"strconv"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// first check  the user is already registered, otherwise negate the action
	userIDparam, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userObj, err := rt.db.UserLogin(userIDparam)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	var userToFollow types.User
	err = json.NewDecoder(r.Body).Decode(&userToFollow)
	if err != nil || userToFollow.Username == userObj.Username {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followsList, err := rt.db.StartFollowing(userObj.Username, userToFollow.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(followsList)
}
