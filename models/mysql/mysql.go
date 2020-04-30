/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         mysql.go
@ Create Time:  2020/4/28 16:09
@ Software:     GoLand
*/

package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-example/models/config"
	"os"
)

var db *gorm.DB

func initDB() (err error) {
	if db, err = gorm.Open("mysql", config.Config.Mysql.Uri); err != nil {
		return
	}

	db.SingularTable(true)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)

	switch os.Getenv("ENV") {
	case "prod":
	default:
		db = db.Debug()
	}
	return
}

func GetConn() (conn *gorm.DB, err error) {
	if db == nil || db.DB().Ping() != nil {
		err = initDB()
	}
	return db, err
}
