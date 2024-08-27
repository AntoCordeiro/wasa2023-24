package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"strconv"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Authentication: check  the user is already registered, otherwise negate the action
	userIDparam, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = rt.db.UserLogin(userIDparam)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//user identified: now get the desired profile
	profileUsername := ps.ByName("profileUsername")
	userProfile, err := rt.db.GetProfile(profileUsername)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} //Have to add the check if the user is banned
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userProfile)
}
