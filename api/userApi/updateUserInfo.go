package userApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
	"regexp"
)

func UpdateUserInfo(c *gin.Context) {
	avatar,avatarIsExist := c.GetPostForm("avatar")
	introduction,introductionIsExist:=c.GetPostForm("introduction")

	if !introductionIsExist || !avatarIsExist {
		api.ReturnJson(c,e.INVALID_PARAMS,e.GetMsg(e.INVALID_PARAMS),"")
		return
	}

	user:= api.GetUser(c)

	// 去除头像URL中可能包含的域名
	reg:=regexp.MustCompile("http.//[^/]+/")
	avatar=reg.ReplaceAllLiteralString(avatar,"/")

	models.UpdateUserInfo(&user, map[string]interface{}{
		"Avatar":avatar,
		"Introduction":introduction,
	})

	//if avatarFile==nil { //只更新介绍
	//	models.UpdateUserInfo(&user, map[string]interface{}{
	//		"Introduction":introduction,
	//	})
	//}else{
	//	//filename:=string( time.Now().Unix() ) + "u" + string( user.ID ) + path.Ext( avatarFile.Filename)
	//	filename:=fmt.Sprintf("%d%s%d%s",time.Now().Unix(),"u",user.ID,path.Ext(avatarFile.Filename))
	//	c.SaveUploadedFile(avatarFile,"public/assets/avatar/"+filename)
	//	avatarUrl:="/assets/avatar/"+filename
	//
	//	models.UpdateUserInfo(&user, map[string]interface{}{
	//		"Avatar":avatarUrl,
	//		"Introduction":introduction,
	//	})
	//}

	newInfo,_:=models.GetUserById(user.ID)

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),newInfo)
}