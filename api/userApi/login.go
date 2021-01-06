package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"androidHomeworkApi/pkg/util"
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context)  {
	name:=c.PostForm("name")
	passwd:=c.PostForm("passwd")

	if name=="" || passwd=="" {
		api.ReturnJson(c,e.ERROR_LOGIN_FAIL,e.GetMsg(e.ERROR_LOGIN_FAIL),"")
		return
	}

	isPass,user  := models.CheckUser(name,passwd)
	if !isPass {
		api.ReturnJson(c,e.ERROR_LOGIN_FAIL,e.GetMsg(e.ERROR_LOGIN_FAIL),"")
		return
	}

	token,err:= util.GenerateToken(user.ID,user.Name)
	if err!=nil {
		api.ReturnJson(c,e.ERROR_AUTH_TOKEN,e.GetMsg(e.ERROR_AUTH_TOKEN),"")
		return
	}

	data:=gin.H{
		"user":user,
		"token":token,
	}

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)
}
