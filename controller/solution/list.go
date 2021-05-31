package solution

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
)

func GetList(c *gin.Context) {
	var solutions []model.Solution
	db := orm.GetDB()
	var problemId = c.Param("problemId")
	var offset = c.DefaultQuery("offset", "0")
	var limit = c.DefaultQuery("limit", "15")
	var total int
	if db.Table("solution").Where("problem_id = ?", problemId).Order("created_at desc").Count(&total).Offset(offset).Limit(limit).
		Find(&solutions).RecordNotFound() {
		c.JSON(404, gin.H{"message": "抱歉，暂无题解"})
		return
	}
	var response []interface{}
	for _, solution := range solutions {
		data := solution.GetData("list")
		response = append(response, data)
	}
	c.JSON(200, gin.H{
		"solutions": response,
		"total":   total,
	})
}
