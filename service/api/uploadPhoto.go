package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/types"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"strconv"
	"io/ioutil"
	"time"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	defer file.Close()

	photoData, err := ioutil.ReadAll(file) 
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	photoObj := types.Photo {
		Username: userObj.Username,
		PhotoData: photoData,
		UploadDate: time.Now(),
		LikesCount: 0,
		CommentsCount: 0,
	}

	err = rt.db.InsertPhoto(photoObj)
	if err!= nil {
		http.Error(w, "Error uploading photo", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photoObj)
}
