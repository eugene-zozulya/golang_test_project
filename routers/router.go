package routers

import (
	"wallet-test/controllers"
	middleware "wallet-test/middleware/validator"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/balance", &controllers.BallanceController{})

	beego.InsertFilter("/balance", beego.BeforeRouter, middleware.BalanceValidator)
}
