package friendApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func AddFriend(c *gin.Context) {
	friendId,err:=com.StrTo( c.PostForm("friend_id") ).Int()

	if friendId<1 || err!=nil {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	// 用户是否存在
	_,isExist := models.GetUserById(friendId)
	if !isExist {
		api.ReturnJson(c,e.ERROR_ADD_FRIEND_NO_EXIST,e.GetMsg(e.ERROR_ADD_FRIEND_NO_EXIST),"")
		return
	}

	user:= api.GetUser(c)

	if user.ID == friendId{
		api.ReturnJson(c,e.ERROR_ADD_FRIEND_SELF,e.GetMsg(e.ERROR_ADD_FRIEND_SELF),"")
		return
	}

	// 是否已经是好友
	isFriend:=models.IsFriend(user.ID,friendId)
	if isFriend {
		api.ReturnJson(c,e.ERROR_ALREADY_FRIEND,e.GetMsg(e.ERROR_ALREADY_FRIEND),"")
		return
	}

	models.AddFriend(models.Friend{
		UserId: user.ID,
		FriendId: friendId,
	})

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),"")
}
