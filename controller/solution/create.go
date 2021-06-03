package solution

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"qkcode/boot/orm"
	"qkcode/model"
	"fmt"
)


type AddSolution struct {
	User  		   string `json:"user" binding:"required"`
	Title          string `json:"title" binding:"required"`
	Content        string `json:"content" binding:"required"`
	Language       string `json:"language" binding:"required"`
}


func Create(c *gin.Context) {
	var data AddSolution
	var id uuid.UUID
	if err := c.BindJSON(&data); err != nil {
		fmt.Print(err)
		c.JSON(422, gin.H{"message": "参数格式错误"})
		return
	}
	db := orm.GetDB()
	db.Table("user").Where("id = ?", data.User).First(&id)
	var solutionId = uuid.NewV4()
	var solution = model.Solution{
		ID:             solutionId,
		Title:          data.Title,
		Content:        data.Content,
		UserId:			id,
	}

	if err := db.Create(&solution).Error; err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "题解提交成功"})

}
