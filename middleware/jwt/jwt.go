package jwt

import (
	"androidHomeworkApi/models"
	"androidHomeworkApi/pkg/e"
	"androidHomeworkApi/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("token")
		var claims *util.Claims
		if token == "" {
			code = e.ERROR_AUTH
		} else {
			var err error
			claims, err = util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		user:=make(map[string]interface{})
		user["ID"]=claims.ID
		user["name"]=claims.Username

		c.Set("user",models.User{
			ID: claims.ID,
			Name: claims.Username,
		})
		c.Next()
	}
}

func JWTNotMust() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token!="" {
			var code int
			var data interface{}
			var claims *util.Claims

			code = e.SUCCESS

			var err error
			claims, err = util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

			if code != e.SUCCESS {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code" : code,
					"msg" : e.GetMsg(code),
					"data" : data,
				})

				c.Abort()
				return
			}

			user:=make(map[string]interface{})
			user["ID"]=claims.ID
			user["name"]=claims.Username

			c.Set("user",models.User{
				ID: claims.ID,
				Name: claims.Username,
			})

		}

		c.Next()
	}
}