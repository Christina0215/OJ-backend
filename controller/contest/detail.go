package contest

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"time"
)

/**
 * @api {POST} contest/ 创建竞赛-Creat
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

func GetDetail(c *gin.Context) {
	var contest model.Contest
	contestId := c.Param("contestId")

	db := orm.GetDB()
	if db.Where("ID = ?", contestId).First(&contest).RecordNotFound() {
		c.JSON(401, gin.H{"message": "抱歉，记录为空"})
		return
	}

	db.Where("ID = ?", contestId).Preload("ContestXProblem").First(&contest)
	data := contest.GetData("detail")
	var response []interface{}
	_user, _ := c.Get("user")
	if _user != nil && (_user.(model.User).RoleID == 1 || time.Now().Unix() >= (data["startAt"]).(int64)) {
		var dataProblem []model.ContestXProblem
		db.Where("contest_id = ?", contestId).Find(&dataProblem)
		for _, data := range dataProblem {
			var problemId = data.ProblemId
			var problem model.Problem
			db.Model(&problem).Where("ID = ?", problemId).First(&problem)
			response = append(response, problem.GetData("detail"))
		}
	}
	c.JSON(200, gin.H{
		"problems": response,
		"data":     data,
	})
}
