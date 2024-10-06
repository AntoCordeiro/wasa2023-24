package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// user
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:myUsername/username", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:myUsername/profiles/:profileUsername", rt.wrap(rt.getUserProfile))

	// stream
	rt.router.GET("/users/:myUsername/stream", rt.wrap(rt.getMyStream))

	// photos
	rt.router.POST("/users/:myUsername/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:myUsername/photos/:photoID", rt.wrap(rt.deletePhoto))

	// follow
	rt.router.POST("/users/:myUsername/follows", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:myUsername/follows/:followedUsername", rt.wrap(rt.unfollowUser))

	// ban
	rt.router.GET("/users/:myUsername/bans", rt.wrap(rt.getBans))
	rt.router.POST("/users/:myUsername/bans", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:myUsername/bans/:bannedUsername", rt.wrap(rt.unbanUser))

	// like
	rt.router.POST("/users/:myUsername/photos/:photoID/likes", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:myUsername/photos/:photoID/likes/:likeID", rt.wrap(rt.unlikePhoto))

	// comment
	rt.router.GET("/users/:myUsername/photos/:photoID/comments", rt.wrap(rt.getComments))
	rt.router.POST("/users/:myUsername/photos/:photoID/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:myUsername/photos/:photoID/comments/:commentID", rt.wrap(rt.uncommentPhoto))

	return rt.router
}
