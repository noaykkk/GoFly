package test

import (
	"CloudStorage/models"
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/cloudstorage?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(dst.String())
}
