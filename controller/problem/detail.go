package problem

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
)

/**
 * @api {GET} problem/:problemId 获取题目-GetOneProblem
 * @apiGroup Problem
 * @apiName GetOneProblem
 * @apiPermission All
 * @apiSuccess {string} title 题目标题
 * @apiSuccess {string} difficulty 题目难度
 * @apiSuccess {string} timeLimit 时间限制
 * @apiSuccess {string} memoryLimit 内存限制
 * @apiSuccess {string} type 题目类型(one problem will own many types, so the string is split by "/"）
 * @apiSuccess {string} content 题目内容
 * @apiSuccess {string} standardInput 标准输入
 * @apiSuccess {string} standardOutput 标准输出
 * @apiSuccess {string} tip 题目标签
 * @apiSuccess {string} samples 题目样例
 * @apiSuccessExample {json} Success-Example:
 * 	{
 *      'title': '一道简单的题目',
 *      'difficulty': '简单',
 *      'timeLimit': '1000ms',
 *      'memoryLimit': '1Mb',
 *      'type': '上机题/C语言'
 *      'content': '似乎这是题目'
 *      'standardInput': '1 2 3'
 *      'standardOutput': '4 5 6'
 *      'tip': '咕咕咕'
 *      'samples': '1 2 3   4 5 6'
 * 	},
 */

func GetDetail(c *gin.Context) {

	var problem model.Problem
	db := orm.GetDB()
	problemId := c.Param("problemId")
	if db.Where("ID = ?", problemId).First(&problem).RecordNotFound() {
		c.JSON(401, gin.H{"message": "该题目不存在哦"})
		return
	}
	c.JSON(200, problem.GetData("detail"))
}
