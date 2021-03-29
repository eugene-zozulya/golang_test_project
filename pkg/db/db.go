package db

import (
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Msql *gorm.DB

func Init() {

	host, _ := beego.AppConfig.String("mysql.host")
	login, _ := beego.AppConfig.String("mysql.login")
	password, _ := beego.AppConfig.String("mysql.password")
	name, _ := beego.AppConfig.String("mysql.name")

	dsn := login + ":" + password + "@tcp(" + host + ":3306)/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	Msql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migrate()
}
