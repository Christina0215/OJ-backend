package solution

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
)

func Delete(c *gin.Context) {
	var solution model.Solution
	db := orm.GetDB()
	ID := c.Param("solutionId")
	db.First(&solution, ID)
	//if problem.ID == 0 {
	//c.JSON(404,gin.H{"message": "该题目不存在"})
	//}
	if err := db.Where("ID = ?", ID).Delete(&solution).Error; err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"message": "原题解删除成功！"})
}
