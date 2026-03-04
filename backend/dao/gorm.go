package dao

import (
	"backend/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {
	conf := config.GetConfig()
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.Username,
		conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database,
	)))
	if err != nil {
		panic(err)
	}
	return db
}
