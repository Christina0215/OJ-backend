package record

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"strconv"
)

/**
 * @api {GET} /problem/:problemId/record/:recordId 获取单个记录-GetOneRecord
 * @apiGroup Record
 * @apiName Detail
 * @apiPermission All
 * @apiSuccess {int} total 测试数据总数
 * @apiSuccess {obj} judgeResult 测试结果
 * @apiSuccess {int} timeCost 消耗时间
 * @apiSuccess {int} memoryCost 消耗内存
 * @apiSuccess {string} compileInfo 测试信息
 * @apiSuccess {obj} language 使用语言
 * @apiSuccess {string} code 提交的代码
 * @apiSuccess {obj} testCase 最后一个错误测试点，若题目测试成功，则返回最后一个成功数据
 * @apiSuccess {float} timeProportion 时间消耗超过百分之多少人
 * @apiSuccess {float} memoryProportion 内存消耗超过百分之多少人
 * @apiSuccessExample {json} Success-Example:
 * 	{
 *      'total': '100',
 *      'judgeResult': {'alias', 'en', 'zh','color'},
 *      'timeCost': '100ms',
 *      'memoryCost': '32MiB',
 *      'compileInfo': '.....'
 *      'language': {'C','c'}
 *      'code': '一段代码'
 *      'testCase': '100'
 *      'timeProportion': '0.700000'
 *      'memoryProportion': '0.310000'
 * 	},
 */

func GetDetail(c *gin.Context) {
	db := orm.GetDB()
	var record model.Record
	var problemId = c.Param("problemId")
	var recordId = c.Param("recordId")
	if db.Where("problem_id = ? AND ID = ?", problemId, recordId).
		First(&record).RecordNotFound() {
		c.JSON(404, gin.H{"message": "抱歉，记录为空"})
		return
	}

	db.Model(&record).Related(&record.JudgeResult).Related(&record.Language).Related(&record.Code).Related(&record.Problem).Related(&record.Code)
	var response = record.GetData("detail")
	if record.JudgeResultId == 1 || record.JudgeResultId == 2 {
		c.JSON(200, response)
	}

	var successRecordNum int
	var lessMemoryNum = 0
	var lessTimeNum = 0
	var testcase model.Testcase
	recordsQuery := db.Table("record").Where("problem_id = ?", problemId).Where("judge_result_id = 3").Not("id = ?", record.ID)
	recordsQuery.Where("time_cost > ?", record.TimeCost).Count(&lessTimeNum)
	recordsQuery.Where("memory_cost > ?", record.MemoryCost).Count(&lessMemoryNum)
	recordsQuery.Count(&successRecordNum)
	if record.JudgeResultId == 3 {
		var memoryProportion, timeProportion float64
		if successRecordNum == 0 {
			memoryProportion = float64(100)
			timeProportion = float64(100)
		} else {
			memoryProportion, _ = strconv.ParseFloat(fmt.Sprintf("%0.4f", float64(lessMemoryNum)/float64(successRecordNum)), 64)
			timeProportion, _ = strconv.ParseFloat(fmt.Sprintf("%0.4f", float64(lessTimeNum)/float64(successRecordNum)), 64)
		}
		response["memoryProportion"] = memoryProportion
		response["timeProportion"] = timeProportion
	} else {
		db.Table("testcase").Where("record_id = ?", record.ID).Not("judge_result_id = 3").First(&testcase)
		response["testcase"] = testcase.GetData()
	}

	c.JSON(200, response)
}
