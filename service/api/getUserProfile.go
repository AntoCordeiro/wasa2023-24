package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Authentication: check  the user is already registered, otherwise negate the action
	_, err := rt.db.UserLogin(ps.ByName("username"))
	if err != nil{
		http.Error(w, "here" + err.Error(), http.StatusInternalServerError)
		return
	}
	//user identified: now get the desired profile
	profileUsername := ps.ByName("profileUsername")
	userProfile, err := rt.db.GetProfile(profileUsername)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userProfile)
}
