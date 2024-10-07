package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// login
	rt.router.POST("/session", rt.wrap(rt.session))

	// follow and unfollow user
	rt.router.PUT("/user/:u_name/followed/:name_to_follow", rt.wrap(rt.followUser))
	rt.router.DELETE("/user/:u_name/followed/:name_to_follow", rt.wrap(rt.unFollowUser))

	// followed and followers list
	rt.router.GET("/user/:u_name/followed_list", rt.wrap(rt.getFollowedList))
	rt.router.GET("/user/:u_name/followers_list", rt.wrap(rt.getFollowersList))

	// ban and unban
	rt.router.PUT("/user/:u_name/ban_user/:user_to_ban", rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:u_name/ban_user/:user_to_ban", rt.wrap(rt.unBanUser))
	rt.router.GET("/user/:u_name/banned_list", rt.wrap(rt.banList))

	// photo upload and delete
	rt.router.POST("/user/:u_name/upload", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/user/:u_name/photo/:photo_id", rt.wrap(rt.deletePhoto))

	// like and unlike
	rt.router.PUT("/user/:u_name/photo/:photo_id/like/:like_name", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/user/:u_name/photo/:photo_id/like/:like_name", rt.wrap(rt.unlikePhoto))

	// comment and delete it
	rt.router.POST("/user/:u_name/photo/:photo_id/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/user/:u_name/photo/:photo_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto))
	// show comments
	rt.router.GET("/user/:u_name/photo/:photo_id/comments", rt.wrap(rt.getComments))

	//  user
	//  get a user profile
	rt.router.GET("/searchuser/:u_name", rt.wrap(rt.getUserProfile))

	//  get a user stream
	rt.router.GET("/user/:u_name/stream", rt.wrap(rt.getStream))

	//  set a new username (only for urself)
	rt.router.PUT("/user/:u_name/set_username", rt.wrap(rt.setMyUserName))

	//  get a list of users
	rt.router.GET("/user/:u_name/searchusers/:user_to_find", rt.wrap(rt.getUsersList))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	rt.router.GET("/photo/:photo_id", rt.wrap(rt.getPhoto))

	return rt.router
}
