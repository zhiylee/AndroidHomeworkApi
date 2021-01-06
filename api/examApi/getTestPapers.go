package examApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTestPapers(c *gin.Context)  {
	testType:=c.Query("type")
	if testType=="" {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	page, err := com.StrTo(c.Query("page")).Int()
	if page<1 || err!=nil {
		page=1
	}

	data := make(map[string]interface{})

	total := models.GetTestPapersTotal(testType)
	paging := api.PagingStatus(total,page)
	api.MergeMap(&data,&paging)

	if page > int( data["pageTotal"].(int) ) {
		api.ReturnJson(c,e.ERROR_NO_TEST_PAPER,e.GetMsg(e.ERROR_NO_TEST_PAPER),data)
		return
	}

	data["lists"],_ = models.GetTestPapers(testType,page)

	api.ReturnJson(c,e.SUCCESS, e.GetMsg(e.SUCCESS),data)
}