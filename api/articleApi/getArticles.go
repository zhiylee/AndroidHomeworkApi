package articleApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetArticles(c *gin.Context)  {
	page, err := com.StrTo(c.Query("page")).Int()
	if page<1 || err!=nil {
		page=1
	}

	data := make(map[string]interface{})
	query := make( map[string]interface{} )

	category,err := com.StrTo(c.Query("category_id")).Int()
	if category>0 && err==nil{
		query["category_id"] = category
	}

	total := models.ArticleTotal(query)
	paging := api.PagingStatus(total,page)
	api.MergeMap(&data,&paging)

	if page > int( data["pageTotal"].(int) ) {
		api.ReturnJson(c,e.TIP_NO_ACTICLE,e.GetMsg(e.TIP_NO_ACTICLE),data)
		return
	}
	
	data["lists"],_ = models.GetArticles(query,page)

	api.ReturnJson(c,e.SUCCESS, e.GetMsg(e.SUCCESS),data)
}