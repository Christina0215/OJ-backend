package record

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
)

/**
 * @api {GET} /problem/:problemId/record 获取记录列表-GetRecordList
 * @apiGroup Record
 * @apiName List
 * @apiPermission All
 * @apiSuccess {array} record[]
 * @apiSuccess {timestamp} createAt 记录创建时间
 * @apiSuccess {obj} language 使用程序语言
 * @apiSuccess {obj} judgeResult 测试结果
 * @apiSuccess {int} timeCost 消耗时间
 * @apiSuccess {int} memoryCost 消耗内存
 * @apiSuccess {string} compileInfo 测试信息
 * @apiSuccessExample {json} Success-Example:
 *{
 *		'record': array,
 *		{
 * 			{
 *      		'createAt': '2019-10-03',
 *      		'language': {'C','c'},
 *      		'judgeResult': '{'alias', 'en', 'zh','color'}',
 *				'timeCost': '100ms',
 *				'memoryCost': '37MiB',
 *				'compileInfo': '...',
 * 			},
 * 			{
 *      		'createAt': '2018-10-03',
 *      		'language': {'C','c'},
 *      		'judgeResult': '{'alias', 'en', 'zh','color'}',
 *				'timeCost': '84ms',
 *				'memoryCost': '31MiB',
 *				'compileInfo': '...',
 * 			}
 *		}
 *}
 */

func GetList(c *gin.Context) {
	_user, _ := c.Get("user")
	var records []model.Record
	db := orm.GetDB()
	var problemId = c.Param("problemId")
	var offset = c.DefaultQuery("offset", "0")
	var limit = c.DefaultQuery("limit", "15")
	var total int
	if db.Table("record").Where("problem_id = ?", problemId).Where("user_id = ?", _user.(model.User).ID).
		Not("judge_result_id", []int{1, 2}).Order("created_at desc").Count(&total).Offset(offset).Limit(limit).
		Find(&records).RecordNotFound() {
		c.JSON(404, gin.H{"message": "抱歉，记录为空"})
		return
	}
	var response []interface{}
	for _, record := range records {
		db.Model(&record).Related(&record.JudgeResult).Related(&record.Language).Related(&record.Problem)
		data := record.GetData("list")
		response = append(response, data)
	}
	c.JSON(200, gin.H{
		"records": response,
		"total":   total,
	})
}
