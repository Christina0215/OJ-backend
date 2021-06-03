package solution

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
)


func GetDetail(c *gin.Context) {

	var solution model.Solution
	db := orm.GetDB()
	solutionId := c.Param("SolutionId")
	if db.Where("ID = ?", solutionId).First(&solution).RecordNotFound() {
		c.JSON(401, gin.H{"message": "该题目不存在哦"})
		return
	}
	c.JSON(200, solution.GetData("detail"))
}
