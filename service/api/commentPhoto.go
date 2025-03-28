package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Authenticate user
	userID, err := GetUserID(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	_, err = rt.db.UserLogin(userID, ps.ByName("myUsername"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Get the photo id from the path
	photoID, err := strconv.Atoi(ps.ByName("photoID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Take the content of the comment in the request body
	var commentObj types.Comment
	err = json.NewDecoder(r.Body).Decode(&commentObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the comment under the photo  and encode the returned comments list in the response
	commentObj.UserID = userID
	commentObj.Date = time.Now()
	commentObj.PhotoID = photoID

	commentObj.ID, err = rt.db.AddComment(commentObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(commentObj)
}
