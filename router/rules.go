package router

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/api/articleApi"
	"androidHomeworkApi/api/examApi"
	"androidHomeworkApi/api/friendApi"
	"androidHomeworkApi/api/userApi"
	"androidHomeworkApi/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func rules(r *gin.Engine) *gin.Engine {
	r.GET("/hello", api.Hello)

	r.Static("/assets","public/assets")

	apiR := r.Group("/api")
	apiR.GET("/articles", articleApi.GetArticles)
	apiR.GET("/categories", articleApi.GetCategories)
	apiR.POST("/register", userApi.Register)
	apiR.POST("/login", userApi.Login)
	apiR.GET("/comments", articleApi.GetComments)
	apiR.GET("/testPapers", examApi.GetTestPapers)
	apiR.GET("/testPaper", examApi.GetTestPaper)

	apiAuthNotMust := r.Group("/api")
	apiAuthNotMust.Use(jwt.JWTNotMust())
	apiAuthNotMust.GET("/article", articleApi.GetArticle)

	apiNeedAuth := r.Group("/api")
	apiNeedAuth.Use(jwt.JWT())
	apiNeedAuth.GET("/userInfo", userApi.GetUserInfo)
	apiNeedAuth.POST("/userInfo", userApi.UpdateUserInfo)
	apiNeedAuth.POST("/uploadAvatar", userApi.UploadAvatar)

	apiNeedAuth.GET("/userFavorites", userApi.GetUserFavorites)
	apiNeedAuth.POST("/userFavorites", userApi.AddFavorite)
	apiNeedAuth.POST("/deleteFavorite", userApi.DeleteFavorite)
	apiNeedAuth.POST("/comment", articleApi.AddComment)

	apiNeedAuth.GET("/friends", friendApi.GetFriends)
	apiNeedAuth.POST("friend", friendApi.AddFriend)
	apiNeedAuth.POST("deleteFriend", friendApi.DeleteFriend)
	apiNeedAuth.POST("/searchUser", userApi.SearchUser)

	apiNeedAuth.GET("/testRecords", examApi.GetTestRecords)
	apiNeedAuth.POST("/testRecord", examApi.AddTestRecord)


	// test api
	apiR.POST("addQuestion",examApi.AddQuestion)

	return r
}