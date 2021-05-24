package problem

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
)

/**
 * @api {POST} /problem/:problemId 添加新问题-Create
 * @apiGroup Problem
 * @apiName ModifyProblem
 * @apiPermission User
 * @apiParam {string} title 题目标题（可选）
 * @apiParam {string} difficulty 题目难度（可选）
 * @apiParam {string} timeLimit 时间限制（可选）
 * @apiParam {string} memoryLimit 内存限制（可选）
 * @apiParam {string} type 题目类型（可选）
 * @apiParam {string} content 题目内容（可选）
 * @apiParam {string} standardInput 标准输入（可选）
 * @apiParam {string} standardOutput 标准输出（可选）
 * @apiParam {string} tip 题目标签（可选）
 * @apiParam {string} samples 题目样例（可选）
 * @apiParam {string} file zip格式（可选）
 * @apiParamExample {json} Request-Example:
 * {
 *      'title': '一道简单的题目',
 *      'difficulty': '简单',
 *      'timeLimit': '1000ms',
 *      'memoryLimit': '1Mb',
 *      'type': '上机题'
 *      'content': '似乎这是题目'
 *      'standardInput': '1 2 3'
 *      'standardOutput': '4 5 6'
 *      'tip': '咕咕咕'
 *      'samples': '1 2 3   4 5 6'
 *      'file': '假装这是zip格式'
 * }
 */

type modifyProblem struct {
	Title          string `json:"title"`
	Type           string `json:"type"`
	Difficulty     string `json:"difficulty"`
	Content        string `json:"content"`
	Samples        string `json:"samples"`
	TimeLimit      string `json:"timeLimit"`
	MemoryLimit    string `json:"memoryLimit"`
	StandardInput  string `json:"standardInput"`
	StandardOutput string `json:"standardOutput"`
	Tip            string `json:"tip"`
}

func Modify(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user == nil || _user.(model.User).RoleID != 1 {
		c.AbortWithStatus(403)
		return
	}
	var data modifyProblem
	ID := c.Param("problemId")
	if err := c.BindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": err.Error()})
	}
	db := orm.GetDB()
	if db.Where("ID = ?", ID).First(&model.Problem{}).RecordNotFound() {
		c.JSON(404, gin.H{"message": "题目走丢了哦"})
		return
	}

	var problem model.Problem
	db.Model(&problem).Where("ID=?", ID).Updates(model.Problem{
		Title:          data.Title,
		Type:           data.Type,
		Difficulty:     data.Difficulty,
		Content:        data.Content,
		Samples:        data.Samples,
		TimeLimit:      data.TimeLimit,
		MemoryLimit:    data.MemoryLimit,
		StandardInput:  data.StandardInput,
		StandardOutput: data.StandardOutput,
		Tip:            data.Tip,
	})
	c.JSON(200, gin.H{"message": "原问题修改成功！"})

}
