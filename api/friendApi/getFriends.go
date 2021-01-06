package friendApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetFriends(c *gin.Context)  {
	user:= api.GetUser(c)

	data,total:=models.GetFriends(user.ID)

	if total<1 {
		api.ReturnJson(c,e.ERROR_NO_FRIEND,e.GetMsg(e.ERROR_NO_FRIEND),"")
		return
	}

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)
}
