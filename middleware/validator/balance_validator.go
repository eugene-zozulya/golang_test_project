package validator

import (
	"encoding/json"
	"wallet-test/models"
	"wallet-test/pkg/response"

	"github.com/beego/beego/v2/server/web/context"
)

var BalanceValidator = func(ctx *context.Context) {

	if request_type := ctx.Request.Method; request_type == "POST" {
		//для POST необходимо провалидировать входящий JSON

		//var request post_balance.PostBallanceRequest
		var request models.Transaction

		input_data := ctx.Input.RequestBody

		err := json.Unmarshal(input_data, &request)
		if err != nil {
			resp := response.Response{
				Success: false,
				Data:    nil,
				Error:   err.Error(),
				Code:    400,
			}

			resp.ServeJSON(ctx)
		}

		validation(ctx, "CurrencyValidator", request.Currency)
		validation(ctx, "TransactionTypeValidator", request.Type)

		ctx.Input.SetData("request", request)

	}

}
