package middleware

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"time"
)

func AuthServiceProvider() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := orm.GetDB()
		token := c.Request.Header.Get("Api_Token")
		var apiToken model.ApiToken
		if !db.Where("token = ?", token).First(&apiToken).RecordNotFound() {
			if apiToken.ExpiredAt.Before(time.Now()) {
				c.AbortWithStatusJSON(401, gin.H{"message": "登陆信息已过期"})
			} else {
				db := orm.GetDB()
				var user model.User
				db.Model(&apiToken).Related(&user)
				c.Set("user", user)
				duration, _ := time.ParseDuration("30m")
				db.Model(&apiToken).Update("expired_at", apiToken.ExpiredAt.Add(duration))
			}
		} else {
			c.AbortWithStatusJSON(401, gin.H{"message": "未登陆"})
		}
		c.Next()
	}
}
