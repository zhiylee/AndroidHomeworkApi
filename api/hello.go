package api

import (
	"androidHomeworkApi/pkg/e"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context)  {
	c.JSON(e.SUCCESS,gin.H{
		"code":200,
		"msg":"Hello world!",
	})
}
