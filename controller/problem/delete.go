package problem

import (
	"github.com/gin-gonic/gin"
	"os"
	"qkcode/boot/orm"
	"qkcode/model"
)

func Delete(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user == nil || _user.(model.User).RoleID != 1 {
		c.AbortWithStatus(403)
		return
	}
	var problem model.Problem
	db := orm.GetDB()
	ID := c.Param("problemId")
	db.First(&problem, ID)
	//if problem.ID == 0 {
	//c.JSON(404,gin.H{"message": "该题目不存在"})
	//}
	if err := db.Where("ID = ?", ID).Delete(&problem).Error; err != nil {
		panic(err)
	}
	err := os.Remove("Public/problem/" + ID)
	if err != nil {
		c.JSON(200, gin.H{"message": "原问题删除成功！"})
	}
}
