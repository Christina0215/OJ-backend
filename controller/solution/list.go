package solution

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"strconv"
)


func GetList(c *gin.Context) {
	var language = c.DefaultQuery("type", "")
	var problemID = c.DefaultQuery("id","")
	var limit, _ = strconv.Atoi(c.DefaultQuery("limit", "15"))
	var offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))
	var solutions []model.Solution
	var total int
	var lists = c.DefaultQuery("list", "normal")

	db := orm.GetDB()
	var result = db
	if language != "" {
		result = result.Where("language LIKE ? and problem_id = ?)", "%"+language+"%",problemID).Find(&model.Solution{})
	}

	var response []interface{}
	if lists == "normal" {
		result.Table("solution").Count(&total).Offset(offset).Limit(limit).Find(&solutions).RecordNotFound()
		for _, solution := range solutions {
			var data = solution.GetData("list")
			response = append(response, data)
		}
	}
	c.JSON(200, gin.H{
		"solutions": response,
		"total":    total,
	})
}
