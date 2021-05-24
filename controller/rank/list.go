package rank

import (
	"github.com/gin-gonic/gin"
	"qkcode/boot/orm"
	"qkcode/model"
	"strconv"
)

func GetList(c *gin.Context) {
	var keyword = c.DefaultQuery("keyword", "")
	//var status = c.DefaultQuery("status", "")
	var limit, _ = strconv.Atoi(c.DefaultQuery("limit", "15"))
	var offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))
	var users []model.User
	var total int
	var lists = c.DefaultQuery("list", "normal")

	db := orm.GetDB()
	var result = db
	if keyword != "" {
		result = db.Where("Username LIKE ?", "%"+keyword+"%").Find(&model.User{})
	}
	var response []interface{}
	if lists == "normal" {
		result.Table("user").Count(&total).Offset(offset).Limit(limit).Order("solved").Find(&users).RecordNotFound()
		for _, user := range users {
			var data = user.GetData("detail")
			var record []model.Record
			var passedRecord []string
			result.Table("record").Where("user_id = ?", data["id"]).Find(&record)
			result.Table("record").Where("user_id = ? and judge_result_id = 3",data["id"]).Find(&passedRecord)
			if(len(record)==0) {
				data["passrate"] = 0
			}else{
				data["passrate"] = float64(len(passedRecord))/float64(len(record)) * 100
			}
			response = append(response, data)
		}
	}
	//if lists == "simplify" {
	//	result.Table("problem").Count(&total).Find(&problems)
	//	for _, problem := range problems {
	//		var data = problem.GetData("simplify_list")
	//		response = append(response, data)
	//	}
	//}
	c.JSON(200, gin.H{
		"users": response,
		"total": total,
	})
}