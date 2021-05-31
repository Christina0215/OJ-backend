package problem

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"strconv"
)

/**
 * @api {GET} problem/ 获取题目列表-GetProblem
 * @apiGroup Problem
 * @apiName GetProblem
 * @apiPermission All
 * @apiParam {int} [limit]
 * @apiParam {int} [offset]
 * @apiParam {string} [Type] 题目类型（可选）--在数据库中以斜杠分开
 * @apiParam {string} difficulty 题目难度（可选）
 * @apiParam {string} keyword 关键词搜索（可选）
 * @apiParam {string} status 题目状态：1 为解题成功，2 为提交但不成功，3 为没提交（可选）
 * @apiSuccess {int} [id]
 * @apiSuccess {string} id 文章id
 * @apiSuccess {string} difficulty 题目难度
 * @apiSuccess {string} title 题目标题
 * @apiSuccess {string} type 题目类型
 * @apiSuccessExample {json} Success-Example:
 *{
 *		'total': 12,
 *		{
 * 			{
 *      		'id': '16',
 *      		'title': '反转链表',
 *      		'type': '数据结构练习',
 *				'difficulty': normal,
 * 			},
 * 			{
 *      		'id': '9',
 *      		'title': '深度优先搜索',
 *      		'type': '算法练习',
 *				'difficulty': normal,
 * 			}
 *		}
 *}
 */
func ContainsString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func GetList(c *gin.Context) {
	var keyword = c.DefaultQuery("keyword", "")
	var status = c.DefaultQuery("status", "")
	var difficulty = c.DefaultQuery("difficulty", "")
	var types = c.DefaultQuery("type", "")
	var limit, _ = strconv.Atoi(c.DefaultQuery("limit", "15"))
	var offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))
	var problems []model.Problem
	var total int
	var lists = c.DefaultQuery("list", "normal")
	_user, _ := c.Get("user")
	var records []model.Record

	db := orm.GetDB()
	var result = db
	if status != ""{
		db.Where("user_id=?",_user.(model.User).ID).Find(&records)

		var problems  []string
		var _problems []string
		if status == "attemped"{
			for _,record := range records {
				if record.JudgeResultId != 3 {
					problems = append(problems, record.ProblemId)
				}else{
					_problems = append(_problems,record.ProblemId)
				}
			}
			for i := 0; i <= len(problems); i++ {
				index:=ContainsString(_problems,problems[i])
				if(index!=-1){
					if(i==len(problems)-1){
						problems = problems[:i]
					}else{
						problems = append(problems[:i],problems[i+1:]...)
						i--
					}
				}
			}
			result = db.Where("id IN (?)",problems).Find(&model.Problem{})
		} else if status == "done"{
			for _,record := range records {
				if record.JudgeResultId == 3{
					problems = append(problems, record.ProblemId)
				}
			}
			result = db.Where("id IN (?)",problems).Find(&model.Problem{})
		}else{
			for _,record := range records {
				problems = append(problems, record.ProblemId)
			}
			result = db.Where("id NOT IN (?)",problems).Find(&model.Problem{})
		}

	}
	if keyword != "" {
		result = result.Where("title LIKE ?", "%"+keyword+"%").Or("type LIKE ?", "%"+keyword+"%").Find(&model.Problem{})
	}
	if difficulty != "" {
		result = result.Where("difficulty = ?", difficulty).Find(&model.Problem{})
	}
	if types != "" {
		result = result.Where("type LIKE ?)", "%"+types+"%").Find(&model.Problem{})
	}

	var response []interface{}
	if lists == "normal" {
		result.Table("problem").Count(&total).Offset(offset).Limit(limit).Find(&problems).RecordNotFound()
		for _, problem := range problems {
			var data = problem.GetData("normal_list")
			var status = 0
			var recordSum = 0
			var passSum = 0
			var acceptance float64 = 0
			var passed = 0
			var attempted = 0
			db.Table("record").Where("user_id = ? and problem_id = ? and judge_result_id = 3",_user.(model.User).ID,data["id"]).Count(&passed)
			db.Table("record").Where("user_id = ? and problem_id = ?",_user.(model.User).ID,data["id"]).Count(&attempted)
			if(passed>0){
				status = 1
			}else if(attempted > 0){
				status = 2
			}
			db.Table("record").Where("problem_id = ?",data["id"]).Count(&recordSum)
			db.Table("record").Where("problem_id = ? and judge_result_id = 3",data["id"]).Count(&passSum)
			data["status"] = status
			if(recordSum == 0 ){
				acceptance = 0
			} else{
				acceptance = float64(passSum)/float64(recordSum)*100
			}
			data["acceptance"] = fmt.Sprintf("%.2f", acceptance)
			response = append(response, data)
		}
	}
	if lists == "simplify" {
		result.Table("problem").Count(&total).Find(&problems)
		for _, problem := range problems {
			var data = problem.GetData("simplify_list")
			response = append(response, data)
		}
	}
	c.JSON(200, gin.H{
		"problems": response,
		"total":    total,
	})
}
