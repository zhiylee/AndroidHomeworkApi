package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"androidHomeworkApi/pkg/util"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context)  {
	name:=c.PostForm("name")
	passwd:=c.PostForm("passwd")

	if name=="" || passwd=="" {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	_,isExist:=models.GetUserByName(name)
	if isExist {
		api.ReturnJson(c,e.ERROR_USER_EXISTS,e.GetMsg(e.ERROR_USER_EXISTS),"")
		return
	}

	passwdHash := util.SHA256Code(passwd)

	data:=models.AddUser(models.User{
		Name: name,
		Passwd: passwdHash,
		Avatar: "/assets/avatar/default_avatar.png",
	})

	//todo auto login after register

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)
}