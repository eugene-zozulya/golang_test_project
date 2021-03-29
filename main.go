package main

import (
	_ "wallet-test/routers"

	"wallet-test/pkg/application_error_handler"
	db "wallet-test/pkg/db"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	defer application_error_handler.Error_handler()

	db.Init()
	beego.Run()
}
