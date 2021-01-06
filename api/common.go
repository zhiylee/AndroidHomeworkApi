package api

import (
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/setting"
	"github.com/gin-gonic/gin"
	"math"
)

func ReturnJson(c *gin.Context,code int, msg string,data interface{})  {
	c.JSON(200,gin.H{
		"code": code,
		//"msg" : e.GetMsg(code),
		"msg" : msg,
		"data" : data,
	})
}

func PagingStatus(total,page int) map[string]interface{} {
	res := make(map[string]interface{})

	res["total"] = total
	res["currentPage"] = page
	res["pageTotal"] = int( math.Ceil( float64(total)/float64(setting.PageSize) ) )

	return res
}

func MergeMap(v1,v2 *map[string]interface{})  {
	for key,val := range *v2 {
		(*v1)[key] = val
	}
}

func GetUser(c *gin.Context) models.User {
	user, isExit :=c.Get("user")

	if !isExit {
		return models.User{}
	}

	return user.(models.User)
}