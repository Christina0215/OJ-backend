package solution

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
)


type modifySolution struct {
	Title          string `json:"title" binding:"required"`
	Content        string `json:"content" binding:"required"`
	Language       string `json:"language" binding:"required"`
}

func Modify(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user == nil || _user.(model.User).RoleID != 1 {
		c.AbortWithStatus(403)
		return
	}
	var data modifySolution
	ID := c.Param("solutionId")
	if err := c.BindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": err.Error()})
	}
	db := orm.GetDB()
	if db.Where("ID = ?", ID).First(&model.Solution{}).RecordNotFound() {
		c.JSON(404, gin.H{"message": "题解走丢了哦"})
		return
	}

	var solution model.Solution
	db.Model(&solution).Where("ID=?", ID).Updates(model.Solution{
		Title:          data.Title,
		Content:        data.Content,
		Language:       data.Language,
	})
	c.JSON(200, gin.H{"message": "原题解修改成功！"})

}
