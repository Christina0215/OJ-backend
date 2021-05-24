package contest

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"qkcode/boot/orm"
	"qkcode/model"
)

/**
 * @api {POST} contest/ 创建竞赛-Create
 * @apiGroup Contest
 * @apiName Creat
 * @apiPermission All
 * @apiParam {string} title
 * @apiParam {string} introduction
 * @apiParam {timestamp} startAt
 * @apiParam {timestamp} endAt
 * @apiParam {bool} enabled
 * @apiParam {array} problem
 * @apiParam {string} problemId
 * @apiParam {order} int
 * @apiParam {int} baseScore
 */

type Problem struct {
	ProblemId string `json:"problem_id" binding:"required"`
	Order     int    `json:"order" binding:"required"`
	BaseScore int    `json:"base_score" binding:"required"`
}

type CreateValidate struct {
	Title        string `json:"title" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Notification string `json:"notification" binding:"required"`

	StartAt int64      `json:"start_at" `
	EndAt   int64      `json:"end_at" `
	Enabled bool       `json:"enabled"`
	Problems []Problem `json:"problems" binding:"required"`
}

func Create(c *gin.Context) {
	//_user, _ := c.Get("user")
	//if _user == nil {
	//	c.JSON(401, gin.H{"message": "登录信息已过期，请重新登录"})
	//	return
	//}
	var data CreateValidate
	var contestId = uuid.NewV4()
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": "格式不正确哦" + err.Error()})
		return
	}
	var contest = model.Contest{
		ID:           contestId,
		Title:        data.Title,
		Introduction: data.Introduction,
		Notification: data.Notification,
		StartAt:      data.StartAt,
		EndAt:        data.EndAt,
	}

	db := orm.GetDB()
	if err := db.Create(&contest).Error; err != nil {
		panic(err)
	}

	for _, problem := range data.Problems {
		var ContestProblem = model.ContestXProblem{
			ContestId: contestId,
			ProblemId: uuid.FromStringOrNil(problem.ProblemId),
			Order:     problem.Order,
			BaseScore: problem.BaseScore,
		}
		if err := db.Create(&ContestProblem).Error; err != nil {
			panic(err)
		}
	}
	c.JSON(200, gin.H{"message": "竞赛创建成功"})
}
