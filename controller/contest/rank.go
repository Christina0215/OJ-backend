package contest
//
//import (
//	"qkcode/boot/orm"
//	_ "qkcode/Controller/Record"
//	_ "qkcode/Controller/Record"
//	"qkcode/model"
//	"github.com/gin-gonic/gin"
//)
//
//func GetRank(c *gin.Context)  {
//	var contestxusers []model.ContestXUser
//	var contestxproblems []model.ContestXProblem
//	var records []model.Record
//	contestId := c.Param("ContestId")
//
//	db := orm.GetDB()
//	if db.Where("ID = ?", contestId).First(&contestxusers).RecordNotFound() {
//		c.JSON(401, gin.H{"message": "抱歉，记录为空"})
//		return
//	}
//
//	var response []interface{}
//	var times int
//	for _, contestxuser := range contestxusers {
//		db.Where("ContestId = ?", contestId).Preload("contest.ContestXProblem").Preload("contest.Record").Find(&contestxuser)
//		var data
//		for _, contestxproblem := range contestxproblems {
//			db.model(&records).Where("ContestId = ?", contestId).Preload("contest.ContestXUser")
//			db.Table("Record").Where("UserId = ?", records.UserId).Where("ProblemId = ?",  contestxproblem.Record.ProblemId).Count(&times)
//			data2["times"] = times
//			data2["score"] = contestxproblem.Record.score
//			db.Table("Record").Where("ProblemId = ?", ProblemId).First(&problem)
//			flag := 0
//			for _, allRecord := range records {
//				db.Table("Record").Where("ProblemId = ?",contestxproblem.Record.ProblemId).Where("judge_result_id = 3")
//				if allRecord.CreatedAt.(string) > records.CreatedAt.(string)  {
//					flag = flag + 1
//				}
//			}
//			if flag != 0 {
//				data2["first"] = true
//			}
//			data2 := problem.UserId
//			if data2["userId"] == UserId{
//				data2["first"] = true
//			}
//			all_score = data2["score"] + all_score
//			response1 = append(response1, data2)
//		}
//		data["allScore"] = all_score
//		response2 = append(response2, data["score"])
//
//	}
//
//}