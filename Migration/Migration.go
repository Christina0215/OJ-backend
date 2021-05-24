package migration

import (
	"fmt"
	"qkcode/boot/orm"
)

var Migration *migration

type migration struct {
	_model map[string]interface{}
}

func InitMigration(model map[string]interface{}) {
	Migration = new(migration)
	Migration._model = model
}

func Fresh() {
	db := orm.GetDB()
	for key, value := range Migration._model {
		if db.HasTable(key) {
			fmt.Println("table", key, "is existed")
		} else {
			db.Table(key).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(value)
			fmt.Println("table", key, "has been created successfully")
		}
	}
}

func Refresh() {
	db := orm.GetDB()
	for key, value := range Migration._model {
		db.DropTableIfExists(key)
		fmt.Println("table", key, "has been dropped successfully")
		db.Table(key).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(value)
		fmt.Println("table", key, "has been created successfully")
	}
}
