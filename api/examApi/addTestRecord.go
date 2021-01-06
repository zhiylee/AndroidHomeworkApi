package examApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"time"
)

func AddTestRecord(c *gin.Context)  {
	testPaperId, err1 := com.StrTo(c.PostForm("test_paper_id")).Int()
	score, err2 := com.StrTo(c.PostForm("score")).Int()
	if testPaperId<1 || err1!=nil || err2!=nil{
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	_,isExist:=models.GetTestPaperById(testPaperId)
	if !isExist {
		api.ReturnJson(c,e.ERROR_NO_EXIST_TEST_PAPER,e.GetMsg(e.ERROR_NO_EXIST_TEST_PAPER),"")
		return
	}

	user:= api.GetUser(c)

	models.AddTestRecord(models.TestRecord{
		TestPaperId: testPaperId,
		UserId: user.ID,
		CreateAt: int( time.Now().Unix() ),
		Score:score,
	})

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),"")
}
