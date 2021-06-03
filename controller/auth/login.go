package auth

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/log"
	"qkcode/boot/orm"
	"qkcode/model"
	"qkcode/utils"
)

/**
 * @api {POST} auth/login 登陆-Login
 * @apiGroup Auth
 * @apiName Login
 * @apiPermission All
 * @apiParam {string} password 密码(sha1加密)
 * @apiParam {string} username 名字
 * @apiParam {boolean} [remember] 记住我(30天)
 * @apiParamExample {json} Request-Example:
 * {
 *      'username': 'administrator',
 *      'password': 'd033e22ae348aeb5660fc2140aec35850c4da997',
 *      'remember': false
 * }
 * @apiParamExample {json} Request-Example2:
 * {
 *      'username': 'administrator',
 *      'password': 'd033e22ae348aeb5660fc2140aec35850c4da997'
 * }
 * @apiSuccess {string} token Api-Token
 * @apiSuccessExample {json} Success-response:
 * {
 *     'token': 'b2336207-3136-47aa-9362-de45f3e49e65'
 * }
 */

type LoginValidate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remember bool   `json:"remember"`
}

func Login(c *gin.Context) {
	var data LoginValidate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(422, gin.H{"err_msg": "格式不正确哦"})
		log.Info(err.Error())
		return
	}
	var user model.User
	db := orm.GetDB()
	if db.Where("username = ?", data.Username).First(&user).RecordNotFound() {
		c.JSON(400, gin.H{"err_msg": "用户名错误"})
		return
	}

	if !utils.Sha256Check(user.Password, data.Password) {
		c.JSON(400, gin.H{"err_msg": "密码错误"})
		return
	}
	var apiToken = model.ApiToken{UserId: user.ID}
	apiToken.AddTime(data.Remember)
	db.Create(&apiToken)
	c.JSON(200, gin.H{"token": apiToken.Token})
}
