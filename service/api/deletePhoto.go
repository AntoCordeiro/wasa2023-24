package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	//"encoding/json"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"strconv"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	photoIDparam, err := strconv.Atoi(ps.ByName("photoID"))
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = rt.db.RemovePhoto(userObj.Username, photoIDparam)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
