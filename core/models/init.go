package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/cloudstorage?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("New Engine Error:%v", err)
		return nil
	}
	return engine
}
