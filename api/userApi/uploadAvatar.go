package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"time"
)

func UploadAvatar(c *gin.Context)  {
	avatarFile,err:=c.FormFile("img")
	if err!=nil && avatarFile!=nil {
		api.ReturnJson(c,e.ERROR_UPDATE_USER_INFO_FAIL,e.GetMsg(e.ERROR_UPDATE_USER_INFO_FAIL),"")
		return
	}

	user:=api.GetUser(c)
	data:=gin.H{}

	filename:=fmt.Sprintf("%d%s%d%s",time.Now().Unix(),"u",user.ID,path.Ext(avatarFile.Filename))
	c.SaveUploadedFile(avatarFile,"public/assets/avatar/"+filename)
	data["avatarUrl"]="/assets/avatar/"+filename

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)

}