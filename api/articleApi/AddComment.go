package articleApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"time"
)

func AddComment(c *gin.Context)  {
	articleId, err := com.StrTo(c.PostForm("article_id")).Int()
	content, isExist := c.GetPostForm("content")

	if articleId<1 || err!=nil || !isExist{
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	user:= api.GetUser(c)
	comment:=models.AddComment(models.Comment{
		ArticleId: articleId,
		UserId: user.ID,
		Content: content,
		CreateAt: int(time.Now().Unix()),
	})

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),comment)
}
