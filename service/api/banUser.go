package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Authenticate user
	userID, err := GetUserID(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	userObj, err := rt.db.UserLogin(userID, ps.ByName("myUsername"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the username of the user to ban from the request body
	var userToBan types.User
	err = json.NewDecoder(r.Body).Decode(&userToBan)
	if err != nil || userToBan.Username == userObj.Username {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userToBan.ID, err = rt.db.GetID(userToBan.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add the user to the banned list and encode the returned banned list in the response
	banList, err := rt.db.AddToBanList(userObj.ID, userToBan.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(banList)
}
