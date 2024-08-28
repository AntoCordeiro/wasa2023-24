package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	//user
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:userID/username", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:userID/profiles/:profileUsername", rt.wrap(rt.getUserProfile))
	
	//photos
	rt.router.POST("/users/:userID/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:userID/photos/:photoID", rt.wrap(rt.deletePhoto))
	
	

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
