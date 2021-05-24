package contest

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"qkcode/boot/orm"
	"qkcode/model"
)

/**
 * @api {POST} contest/:contestId 修改竞赛-Modify
 * @apiGroup Contest
 * @apiName Creat
 * @apiPermission All
 * @apiParam {string} title optional
 * @apiParam {string} introduction optional
 * @apiParam {timestamp} startAt optional
 * @apiParam {timestamp} endAt optional
 * @apiParam {bool} enabled optional
 * @apiParam {array} problem optional
 * @apiParam {string} problemId optional
 * @apiParam {order} int optional
 * @apiParam {int} baseScore optional
 */

type modifyValidate struct {
	Title        string `json:"title" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Notification string `json:"notification" binding:"required"`

	StartAt int64      `json:"start_at" `
	EndAt   int64      `json:"end_at" `
	Enabled bool       `json:"enabled"`
	Problems []Problem `json:"problems" binding:"required"`
}

func Modify(c *gin.Context) {
	var contest model.Contest
	var ContestProblem model.ContestXProblem
	_user, _ := c.Get("user")
	if _user == nil {
		c.JSON(401, gin.H{"message": "登录信息已过期，请重新登录"})
		return
	}

	var data modifyValidate
	contestId := c.Param("contestId")
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": "格式不正确哦" + err.Error()})
		return
	}
	db := orm.GetDB()
	if db.Where("ID = ?", contestId).First(&model.Contest{}).RecordNotFound() {
		c.JSON(404, gin.H{"message": "比赛走丢了哦"})
		return
	}
	db.Model(&contest).Where("ID=?", contestId).Updates(model.Contest{
		Title:        data.Title,
		Introduction: data.Introduction,
		Notification: data.Notification,
		StartAt:      data.StartAt,
		EndAt:        data.EndAt,
	});
	db.Where("contest_id=?", contestId).Unscoped().Delete(&ContestProblem)
	for _, problem := range data.Problems {
		var ContestProblem = model.ContestXProblem{
			ContestId: uuid.FromStringOrNil(contestId),
			ProblemId: uuid.FromStringOrNil(problem.ProblemId),
			Order:     problem.Order,
			BaseScore: problem.BaseScore,
		}
		if err := db.Create(&ContestProblem).Error; err != nil {
			panic(err)
		}
	}
	c.JSON(200, gin.H{"message": "contest修改成功"})
}
