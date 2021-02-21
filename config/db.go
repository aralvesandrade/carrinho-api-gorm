package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//InitDb ...
func InitDb(user, password, host, dbName string, port int) (*gorm.DB, error) {
	connect := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, error := gorm.Open("mysql", connect)

	if error != nil {
		return nil, error
	}

	fmt.Println("Database MySQL connected")

	db.LogMode(true)

	//db.Debug().AutoMigrate(&model.Carrinho{})

	return db, nil
}
