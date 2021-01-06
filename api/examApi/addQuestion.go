package examApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"regexp"
)

func AddQuestion(c *gin.Context)  {
	testPaperId,_:=com.StrTo( c.PostForm("test_paper_id") ).Int()
	question,isExist:=c.GetPostForm("question")

	var title,answer string
	options:=make(map[string]string)

	if !isExist {
		title=c.PostForm("title")
		options=c.PostFormMap("options")
		answer=c.PostForm("answer")
	}else{

		reg:= regexp.MustCompile("(.+) A.(.+) B.(.+) C.(.+) D.(.+)[.\n]*参考答案：([A-Z])")
		match:=reg.FindSubmatch([]byte(question))
		title= string(match[1])
		options["A"]=string(match[2])
		options["B"]=string(match[3])
		options["C"]=string(match[4])
		options["D"]=string(match[5])
		answer=string(match[6])

	}



	//optionJson,_ := json.Marshal(options)

	models.AddQuestion(models.Question{
		Title: title,
		Options: options,
		Answer: answer,
		TestPaperId: testPaperId,
	})

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),"")
}