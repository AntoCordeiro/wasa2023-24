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
	rt.router.PUT("/users/:myUsername/username", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:myUsername/profiles/:profileUsername", rt.wrap(rt.getUserProfile))

	//photos
	rt.router.POST("/users/:myUsername/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:myUsername/photos/:photoID", rt.wrap(rt.deletePhoto))

	//follow
	rt.router.GET("/users/:myUsername/follows", rt.wrap(rt.getFollows))
	rt.router.PUT("/users/:myUsername/follows", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:myUsername/follows/:followID", rt.wrap(rt.unfollowUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
