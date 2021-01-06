package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"time"
)

func AddFavorite(c *gin.Context)  {
	articleId, err := com.StrTo(c.PostForm("article_id")).Int()
	if articleId<1 || err!=nil {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	user:= api.GetUser(c)

	isExist,_:=models.IsFavorite(user.ID,articleId)

	if isExist {
		api.ReturnJson(c,e.ERROR_FAVORITE_EXSIT,e.GetMsg(e.ERROR_FAVORITE_EXSIT),"")
		return
	}

	models.AddFavorite(models.Favorite{
		UserId: user.ID,
		ArticleId: articleId,
		CreateAt: int(time.Now().Unix()),
	})

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),"")
}
