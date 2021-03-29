package controllers

import (
	"fmt"
	"strconv"
	"wallet-test/models"
	"wallet-test/pkg/balance"
	"wallet-test/pkg/db"
	"wallet-test/pkg/response"

	beego "github.com/beego/beego/v2/server/web"
)

type PostRequest struct {
	UserID     string  `json:"user_id"`
	Currency   string  `json:"currency"`
	Amount     float64 `json:"amount"`
	TimePlaced string  `json:"time_placed"`
	Type       string  `json:"deposit"`
}

type BallanceController struct {
	beego.Controller
}

func (c *BallanceController) Get() {

	ids := make([]string, 0)
	c.Ctx.Input.Bind(&ids, "id")

	var balances []models.Balance

	rec := db.Msql.Model(balance.Balance{})

	if len(ids) > 0 {
		rec = rec.Where("user_id IN ?", ids)
	}

	type UserData struct {
		UserID   string
		Ballance string
	}
	data := make(map[string][]UserData, 0)

	rec.Find(&balances).Group("currency")

	for _, v := range balances {

		u := UserData{
			UserID:   fmt.Sprint(v.UserID),
			Ballance: strconv.FormatFloat(v.Value, 'f', 3, 64),
		}
		data[v.Currency] = append(data[v.Currency], u)
	}

	resp := response.Response{
		Success: true,
		Data:    data,
		Code:    200,
	}

	resp.ServeJSON(c.Ctx)
}

func (c *BallanceController) Post() {

	transaction := c.Ctx.Input.GetData("request").(models.Transaction)

	var user models.User

	result := db.Msql.Model(&user).First(&user, transaction.UserID)

	if (result.RowsAffected == 0) || (result.Error != nil) {
		resp := response.Response{
			Success: false,
			Data:    nil,
			Error:   "User not found",
			Code:    404,
		}

		resp.ServeJSON(c.Ctx)
	}

	balance := new(balance.Balance)

	balance.Init(user.ID, transaction.Currency)

	if !balance.Get() {
		resp := response.Response{
			Success: false,
			Error:   "Currency not found for this user",
			Code:    404,
		}

		resp.ServeJSON(c.Ctx)
	}

	var err error

	switch transaction.Type {
	case "withdrowal":
		err = balance.Withdrowal(transaction.Amount)
	case "deposit":
		err = balance.Deposit(transaction.Amount)
	}

	if err != nil {
		resp := response.Response{
			Success: false,
			Error:   err.Error(),
			Code:    400,
		}

		resp.ServeJSON(c.Ctx)
	}

	//если дошли сюда, то в переменной balance значение балланса изменилось

	err = balance.Save(transaction)
	if err != nil {
		resp := response.Response{
			Success: false,
			Error:   err.Error(),
			Code:    400,
		}

		resp.ServeJSON(c.Ctx)
	}

	resp := response.Response{
		Success: true,
		Code:    200,
	}

	resp.ServeJSON(c.Ctx)
}
