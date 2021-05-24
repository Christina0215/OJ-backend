package contest

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	uuid "github.com/satori/go.uuid"
)

/**
 * @api {POST} contest/:contestId/sign 参加竞赛-signUp
 * @apiGroup Contest
 * @apiName SignUp
 */

func SignUp(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user == nil {
		c.JSON(401, gin.H{"message": "登录信息已过期，请重新登录"})
		return
	}

	contestId := c.Param("contestId")
	ContestId := uuid.FromStringOrNil(contestId)

	_user, _ = c.Get("user")
	if _user == nil {
		c.JSON(401, gin.H{"message": "登录信息已过期，请重新登录"})
		return
	}
	user := _user.(model.User)
	db := orm.GetDB()
	var role model.Role
	if err := db.Model(user).Related(&role).Find(&role).Error; err != nil {
		panic(err)
	}
	userId := user.ID

	if !db.Where("contestId=? AND userId=?", ContestId, userId).First(&model.ContestXUser{}).RecordNotFound() {
		c.JSON(403, gin.H{"message": "您已经加入本竞赛"})
		return
	}
	var contestXUser = model.ContestXUser{
		ContestId: ContestId,
		UserId:    userId,
	}
	if err := db.Create(&contestXUser).Error; err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"message": "加入竞赛成功"})
}
