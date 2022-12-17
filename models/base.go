package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go_gin_todo/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

var err error

const (
	tableNameTodo = "todos"
)

func init() {

	Db, err = gorm.Open(sqlite.Open(config.Config.DbName))

	if err != nil {
		fmt.Errorf("error:%v", err)
	}
	Db.AutoMigrate(&Todo{})

}
