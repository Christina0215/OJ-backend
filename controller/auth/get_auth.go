package auth

import (
	"github.com/gin-gonic/gin"
	"qkcode/model"
)

/**
 * @api {GET} auth/auth 获取用户信息-GetAuth
 * @apiGroup Auth
 * @apiName GetAuth
 * @apiPermission All
 * @apiSuccess {string} username 用户名
 * @apiSuccess {string} id
 * @apiSuccess {string} avatar 头像
 * @apiSuccess {string} email 邮箱
 * @apiSuccess {integer} gender 性别
 * @apiSuccess {json} role 角色
 * @apiSuccessExample {json} Success-Example:
 * {
 *      'username': 'test',
 *      'id': '22a52817-bc97-4e75-b8cd-a1b5e91cda2f',
 *      'avatar': 'picture.png',
 *      'email': 'haha@example.com'
 *      'gender': 1
 *      'role': {
 *          'roleId': 1,
 *          'alias': 'admin',
 *          'name': '管理员',
 *      }
 * }
 */

func GetAuth(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user == nil {
		c.JSON(401, gin.H{"message": "登录信息已过期，请重新登录"})
		return
	}
	user := _user.(model.User)
	c.JSON(200, user.GetData("detail"))
	return
}
