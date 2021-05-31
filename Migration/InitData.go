package migration

import (
	uuid "github.com/satori/go.uuid"
	"qkcode/boot/orm"
	"qkcode/model"
)

func InitData() {
	db := orm.GetDB()
	var roles = [2]model.Role{
		{ID: 1, Alias: "admin", Name: "管理员"},
		{ID: 2, Alias: "user", Name: "用户"},
	}
	for _, role := range roles {
		db.Table("role").Create(role)
	}

	var languages = [2]model.Language{
		{DisplayName: "C", Extension: "c"},
		{DisplayName: "C++", Extension: "cc"},
	}
	for _, language := range languages {
		db.Table("language").Create(&language)
	}

	var users = [2]model.User{
		{
			ID: uuid.NewV1(),
			RoleID: 1,
			Email: "1532706870@qq.com",
			Gender: true,
			Username: "admin",
			Password: "ca9e680399decb9dd10d0cc4acda282c05e905174ab331bd9503e9f2e3b59f07",
		},
		{
			ID: uuid.NewV1(),
			RoleID: 2,
			Email: "1532706870@qq.com",
			Gender: false,
			Username: "root",
			Password: "ca9e680399decb9dd10d0cc4acda282c05e905174ab331bd9503e9f2e3b59f07",
		},
	}
	for _, user := range users {
		db.Table("user").Create(&user)
	}
	var judgeResults = [11]model.JudgeResult{
		{Alias: "Pending", En: "Pending", Zh: "等待测评", Color: "#9e9e9e"},
		{Alias: "Judging", En: "Judging", Zh: "正在测评", Color: "#2196f3"},
		{Alias: "AC", En: "Accepted", Zh: "通过测试", Color: "#4caf50"},
		{Alias: "PE", En: "Presentation Error", Zh: "输出格式错误", Color: "#ff9800"},
		{Alias: "WA", En: "Wrong Answer", Zh: "错误答案", Color: "#f44336"},
		{Alias: "OLE", En: "Output Limit Exceeded", Zh: "超出输出限制", Color: "#e91e63"},
		{Alias: "TLE", En: "Time Limit Exceeded", Zh: "超出时间限制", Color: "#9c27b0"},
		{Alias: "MLE", En: "Memory Limit Exceeded", Zh: "超出内存限制", Color: "#673ab7"},
		{Alias: "CE", En: "Compilation Error", Zh: "编译错误", Color: "#ffeb3b"},
		{Alias: "RE", En: "Runtime Error", Zh: "运行时错误", Color: "#ff5722"},
		{Alias: "SE", En: "System Error", Zh: "系统错误", Color: "#000000"},
	}
	for _, judgeResult := range judgeResults {
		db.Table("judge_result").Create(&judgeResult)
	}
}
