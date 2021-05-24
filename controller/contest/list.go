package contest

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"strconv"
)

/**
 * @api {GET} contest/ 获取竞赛信息-List
 * @apiGroup Contest
 * @apiName List
 * @apiPermission All
 * @apiSuccess {array} contest
 * @apiSuccess {string} contest.title
 * @apiSuccess {timestamp} contest.startAt
 * @apiSuccess {timestamp} contest.endAt
 * @apiSuccess {bool} enabled
 * @apiSuccess {bool} joined
 */

func GetList(c *gin.Context) {
	var keyword = c.DefaultQuery("keyword", "")
	//var status = c.DefaultQuery("status", "")
	var limit, _ = strconv.Atoi(c.DefaultQuery("limit", "15"))
	var offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))
	var contests []model.Contest
	var total int

	db := orm.GetDB()
	var result = db

	if keyword != "" {
		result = db.Where("title LIKE ?", "%"+keyword+"%").Or("type LIKE ?", "%"+keyword+"%").Find(&model.Problem{})
	}
	var response []interface{}
	result.Table("contest").Count(&total).Offset(offset).Limit(limit).Find(&contests).RecordNotFound()
	for _, contest := range contests {
		var data = contest.GetData("list")
		response = append(response, data)
	}
	c.JSON(200, gin.H{
		"contests": response,
		"total":    total,
	})
}
