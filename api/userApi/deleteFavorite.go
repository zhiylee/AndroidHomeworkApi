package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func DeleteFavorite(c *gin.Context)  {
	//c.Request.ParseForm()
	//a:= c.Params.ByName("article_id")
	//returnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),a)
	//return
	//c.Request.FormValue()
	//c.Reques
	articleId, err := com.StrTo(c.PostForm("article_id")).Int()
	//articleId, err := com.StrTo(c.Request.PostForm("article_id")).Int()
	//articleId, err := com.StrTo(c.Params.ByName("article_id")).Int()
	if articleId<1 || err!=nil {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	user:= api.GetUser(c)

	isExist,id:=models.IsFavorite(user.ID,articleId)

	if !isExist {
		api.ReturnJson(c,e.ERROR_FAVORITE_NO_EXIT,e.GetMsg(e.ERROR_FAVORITE_NO_EXIT),"")
		return
	}

	models.DeleteFavorite(id)

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),"")
}
