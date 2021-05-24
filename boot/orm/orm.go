package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"qkcode/boot/config"
)

var db *gorm.DB

func InitOrm() {
	engine := config.GetString("database.engine")
	dbEngine, err := gorm.Open(engine, getParams(engine))
	if err != nil {
		panic(fmt.Errorf("Fatal error open database error [err=%s]\n", err))
	}
	db = dbEngine
	db.SingularTable(true) //禁止表名复数
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}

func EndOrm() {
	err := db.Close()
	if err != nil {
		panic(fmt.Errorf("Fatal error close database error [err=%s]\n", err))
	}
}

func getParams(engine string) string {
	switch engine {
	case "mysql":
		host := config.GetStringWithDefault("database.host", "localhost")
		port := config.GetStringWithDefault("database.port", "8888")
		dbname := config.GetStringWithDefault("database.dbname", "QKCODE")
		username := config.GetStringWithDefault("database.user", "root")
		password := config.GetStringWithDefault("database.password", "")
		mysqlParams := config.GetStringWithDefault(
			"database.mysqlParams",
			"parseTime=True&charset=utf8mb4&loc=Local")
		params := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
			username, password, host, port, dbname, mysqlParams)
		fmt.Println(params)
		return params
	case "sqlite3":
		params := config.GetStringWithDefault("database.dbname", "/tmp/gorm.db")
		return params
	case "postgres":
		host := config.GetStringWithDefault("database.host", "localhost")
		port := config.GetStringWithDefault("database.port", "8888")
		dbname := config.GetStringWithDefault("database.dbname", "QKCODE")
		username := config.GetStringWithDefault("database.user", "root")
		password := config.GetStringWithDefault("database.password", "")
		sslMode := config.GetStringWithDefault("database.sslmode", "disable")
		params := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, username, password, dbname, sslMode)
		return params
	default:
		panic(fmt.Errorf("Fatal error getting database params: %s\n", engine))
	}
}
