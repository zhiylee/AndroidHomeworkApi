package articleApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetComments(c *gin.Context)  {
	articleId, err := com.StrTo(c.Query("article_id")).Int()
	if articleId<1 || err!=nil {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	page, err := com.StrTo(c.Query("page")).Int()
	if page<1 || err!=nil {
		page=1
	}

	data := make(map[string]interface{})

	total := models.GetArticleCommentsTotal(articleId)
	paging := api.PagingStatus(total,page)
	api.MergeMap(&data,&paging)

	if page > int( data["pageTotal"].(int) ) {
		api.ReturnJson(c,e.ERROR_NO_COMMENT,e.GetMsg(e.ERROR_NO_COMMENT),data)
		return
	}

	data["lists"],_ = models.GetArticleComments(articleId,page)

	api.ReturnJson(c,e.SUCCESS, e.GetMsg(e.SUCCESS),data)

}
