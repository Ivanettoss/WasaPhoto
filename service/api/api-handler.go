package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	//login
	rt.router.POST("/session", rt.wrap(rt.session))

	//follow and unfollow user
	rt.router.PUT("/user/:u_name/followed/:name_to_follow", rt.wrap(rt.followUser))
	rt.router.DELETE("/user/:u_name/followed/:name_to_follow", rt.wrap(rt.unFollowUser))

	//followed and followers list
	rt.router.GET("/user/:u_name/followed_list", rt.wrap(rt.getFollowedList))
	rt.router.GET("/user/:u_name/followers_list", rt.wrap(rt.getFollowersList))

	//ban and unban
	rt.router.PUT("/user/:u_name/ban_user/:user_to_ban", rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:u_name/ban_user/:user_to_ban", rt.wrap(rt.unBanUser))
	rt.router.GET("/user/:u_name/banned_list", rt.wrap(rt.banList))

	//photo upload and delete
	rt.router.POST("/user/:u_name/upload", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/user/:u_name/photo/:photo_id", rt.wrap(rt.deletePhoto))

	/*
		//like and unlike
		rt.router.PUT("/user/:u_name/photo/:photo_id/like/:like_name", rt.wrap(rt.likePhoto))
		rt.router.DELETE("/user/:u_name/photo/:photo_id/like/:like_name", rt.wrap(rt.unlikePhoto))

		//comment and delete it
		rt.router.POST("/user/:u_name/photo/:photo_id/comment", rt.wrap(rt.commentPhoto))
		rt.router.DELETE("/user/:u_name/photo/:photo_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto))

		//show comments
		rt.router.POST("/user/:u_name/photo/:photo_id/comments", rt.wrap(rt.getComments))

		//photo upload and delete
		rt.router.POST("/user/:u_name/upload", rt.wrap(rt.uploadPhoto))
		rt.router.DELETE("/user/:u_name/photo/:photo_id", rt.wrap(rt.deletePhoto))

		//user
		rt.router.GET("/user/:u_name/", rt.wrap(rt.getUserProfile))
	*/

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
