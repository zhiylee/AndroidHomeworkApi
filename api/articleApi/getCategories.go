package articleApi

import (
	"androidHomeworkApi/api"
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context)  {

	data := make(map[string]interface{})
	data["lists"],data["total"] = models.GetCategories()

	api.ReturnJson(c,e.SUCCESS,e.GetMsg(e.SUCCESS),data)
}