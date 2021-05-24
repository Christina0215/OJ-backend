package auth

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"time"
)

/**
 * @api {POST} auth/code 获取验证码-GetCode
 * @apiGroup Auth
 * @apiName GetCode
 * @apiPermission All
 * @apiParam {string} email 验证邮箱
 * @apiParamExample {json} Request-Example:
 * {
 *      'email': 'haha@example.com'
 * }
 * @apiError {string} failed 获取失败
 * @apiErrorExample {json} 403:
 * {
 *      'message': '请不要频繁发送信息哦'
 * }
 */

type EmailValidate struct {
	Email string `json:"email" binding:"required,email"`
}

func GetCode(c *gin.Context) {
	var data EmailValidate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": "邮件格式不正确哦" + err.Error()})
		return
	}
	db := orm.GetDB()
	var pre model.VerifyCode
	if !db.Where("email = ?", data.Email).Last(&pre).RecordNotFound() {
		if !pre.ExpiredAt.Before(time.Now()) {
			c.JSON(403, gin.H{"message": "请不要频繁发送信息"})
			return
		}
	}
	var code = model.VerifyCode{Email: data.Email}
	if err := db.Create(&code).Error; err != nil {
		panic(err)
	}
	message := "你的验证码为 " + code.Code + " 有效期为3分钟"
	code.SendMsg(data.Email, message)
}
