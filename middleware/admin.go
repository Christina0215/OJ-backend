package middleware

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	_ "time"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := orm.GetDB()
		token := c.Request.Header.Get("Api_Token")
		var apiToken model.ApiToken
		if !db.Where("token=?", token).First(&apiToken).RecordNotFound() {
			var role model.Role
			db.Where("ID = ?", apiToken.User.RoleID).First(&role)
			if role.Alias == "1" {
				c.Next()
			} else {
				c.AbortWithStatus(401)
			}

		}
	}
}
