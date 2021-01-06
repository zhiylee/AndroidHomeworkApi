package friendApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func DeleteFriend(c *gin.Context)  {
	friendId,err:=com.StrTo( c.PostForm("friend_id") ).Int()

	if friendId<1 || err!=nil {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	user:= api.GetUser(c)

	models.DeleteFriend(models.Friend{
		UserId:user.ID,
		FriendId: friendId,
	})

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),"")
}
