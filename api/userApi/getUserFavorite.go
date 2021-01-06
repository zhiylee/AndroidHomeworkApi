package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetUserFavorites(c *gin.Context)  {
	page, err := com.StrTo(c.Query("page")).Int()
	if page<1 || err!=nil {
		page=1
	}

	user:= api.GetUser(c)

	data := make(map[string]interface{})

	total := models.GetUserFavoriteTotal(user.ID)
	paging := api.PagingStatus(total,page)
	api.MergeMap(&data,&paging)

	if page > int( data["pageTotal"].(int) ) {
		api.ReturnJson(c,e.ERROR_NO_FAVORITE,e.GetMsg(e.ERROR_NO_FAVORITE),data)
		return
	}

	data["lists"],_ = models.GetUserFavorite(user.ID,page)

	api.ReturnJson(c,e.SUCCESS, e.GetMsg(e.SUCCESS),data)


}
