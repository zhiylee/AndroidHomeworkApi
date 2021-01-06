package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
)

func SearchUser(c *gin.Context)  {
	name,isExist:=c.GetPostForm("name")

	if !isExist {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	user,isExist:=models.GetUserByName(name)
	if !isExist {
		api.ReturnJson(c,e.ERROR_USER_NO_EXIST,e.GetMsg(e.ERROR_USER_NO_EXIST),"")
		return
	}

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),user)

}
