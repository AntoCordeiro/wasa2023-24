package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Authenticate user
	userID, err := GetUserID(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	userObj, err := rt.db.UserLogin(userID, ps.ByName("myUsername"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Get the username of the user to follow from the request body
	var userToFollow types.User
	err = json.NewDecoder(r.Body).Decode(&userToFollow)
	if err != nil || userToFollow.Username == userObj.Username {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the ID of the user to follow from the database
	userToFollow.ID, err = rt.db.GetID(userToFollow.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Insert the follow in the database and encode the returned follows list in the response
	err = rt.db.StartFollowing(userObj.ID, userToFollow.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
