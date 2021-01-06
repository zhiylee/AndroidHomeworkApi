package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context)  {
	user:= api.GetUser(c)

	data,_:=models.GetUserById(user.ID)

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)
}