package record

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"os"
	"qkcode/boot/orm"
	"qkcode/model"
)

/**
 * @api {POST} /problem/:problemId/record 添加新记录-Create
 * @apiGroup Record
 * @apiName Create
 * @apiPermission User
 * @apiParam {string} code 提交的代码
 * @apiParam {string} language 所用代码语言
 * @apiParamExample {json} Request-Example:
 * {
 *      'code': '一段代码'
 *      'language: 'c'
 * }
 */

type AddRecord struct {
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

func Create(c *gin.Context) {
	var data AddRecord
	if err := c.BindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": "参数格式错误"})
		return
	}
	db := orm.GetDB()
	problemId := c.Param("problemId")
	filename := uuid.NewV4().String()
	mimeType := "txt"
	file, err := os.Create("public/code/" + filename + "." + mimeType)
	if err != nil {
		panic(err)
	} else {
		content := data.Code
		_, err = file.Write([]byte(content))
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	_user, _ := c.Get("user")
	var language model.Language
	if err := db.Where("display_name = ?", data.Language).Find(&language).Error; err != nil {
		panic(err)
	}
	var record = model.Record{
		UserId:     _user.(model.User).ID,
		LanguageId: language.ID,
		ProblemId:  problemId,
	}
	if err := db.Create(&record).Error; err != nil {
		panic(err)
	}
	var code = model.Code{
		RecordId:   int(record.ID),
		Filename:   filename + "." + mimeType,
		LanguageId: language.ID,
	}
	if err := db.Create(&code).Error; err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"recordId": record.ID})
}
