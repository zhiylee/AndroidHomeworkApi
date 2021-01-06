package examApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTestPaper(c *gin.Context)  {
	testType:=c.Query("type")
	id,_ := com.StrTo( c.Query("id") ).Int()

	if ( testType=="" && id<1 ) || ( testType!="" && id>0 ) {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	data:=make(map[string]interface{})
	testPaper:=models.TestPaper{}
	var isExist bool

	if id>0 {
		testPaper,isExist = models.GetTestPaperById(id)
	}else{
		testPaper,isExist = models.GetTestPaperByType(testType)
	}

	if !isExist {
		api.ReturnJson(c,e.ERROR_NO_EXIST_TEST_PAPER,e.GetMsg(e.ERROR_NO_EXIST_TEST_PAPER),"")
		return
	}

	data["id"]=testPaper.ID
	data["title"]= testPaper.Title
	data["type"]= testPaper.Type

	data["questions"],_=models.GetQuestions(testPaper.ID)

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)

}