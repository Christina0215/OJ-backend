package problem

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"qkcode/boot/orm"
	"qkcode/model"
	"qkcode/utils"
)

/**
 * @api {POST} problem/ 添加新问题-Create
 * @apiGroup Problem
 * @apiName Create
 * @apiPermission User
 * @apiParam {string} title 题目标题
 * @apiParam {string} difficulty 题目难度
 * @apiParam {string} timeLimit 时间限制
 * @apiParam {string} memoryLimit 内存限制
 * @apiParam {string} type 题目类型
 * @apiParam {string} content 题目内容
 * @apiParam {string} standardInput 标准输入
 * @apiParam {string} standardOutput 标准输出
 * @apiParam {string} tip 题目标签
 * @apiParam {string} samples 题目样例
 * @apiParam {string} file zip格式
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

type AddProblem struct {
	Title          string `json:"title" binding:"required"`
	Type           string `json:"type" binding:"required"`
	Difficulty     string `json:"difficulty" binding:"required"`
	Content        string `json:"content" binding:"required"`
	Samples        string `json:"samples" binding:"required"`
	TimeLimit      string `json:"timeLimit" binding:"required"`
	MemoryLimit    string `json:"memoryLimit" binding:"required"`
	StandardInput  string `json:"standardInput" binding:"required"`
	StandardOutput string `json:"standardOutput" binding:"required"`
	Tip            string `json:"tip"`
	File           string `json:"file" bind:"required"`
}

func Create(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user == nil || _user.(model.User).RoleID != 1 {
		c.AbortWithStatus(403)
		return
	}
	var data AddProblem
	if err := c.BindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": "参数格式错误"})
		return
	}
	db := orm.GetDB()
	if !db.Where("title=?", data.Title).First(&model.Problem{}).RecordNotFound() {
		c.JSON(403, gin.H{"message": "该题目名已存在哇"})
		return
	}

	var problemId = uuid.NewV4()
	var testdataNumber int
	var err error
	if err, testdataNumber = utils.Unzip("public/temp/"+data.File, "public/problem/", problemId.String()); err != nil {
		panic(err)
	}
	var problem = model.Problem{
		ID:             problemId,
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
		UserId:         _user.(model.User).ID,
		TestdataNumber: testdataNumber,
	}

	if err := db.Create(&problem).Error; err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "题目添加成功"})

}
