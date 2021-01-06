package articleApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetArticle(c *gin.Context)  {
	id, err := com.StrTo(c.Query("id")).Int()
	if id<1 || err!=nil {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	query := make( map[string]interface{} )
	query["ID"]=id

	data := models.GetArticle(query)
	if data.ID<1 {
		api.ReturnJson(c,e.TIP_NO_ACTICLE,e.GetMsg(e.TIP_NO_ACTICLE),"")
		return
	}

	// 文章是否收藏
	data.IsFavotite = false // 默认不收藏
	user := api.GetUser(c) //获取用户
	if user.ID<1 {
		api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)
		return
	}


	data.IsFavotite,_ = models.IsFavorite(user.ID,id)


	api.ReturnJson(c,e.SUCCESS, e.GetMsg(e.SUCCESS),data)
}