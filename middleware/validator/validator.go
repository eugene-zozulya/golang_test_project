package validator

import (
	"reflect"
	"wallet-test/pkg/response"
	"wallet-test/pkg/validator"

	"github.com/beego/beego/v2/server/web/context"
)

var Validators map[string]validator.ValidatorInterface

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func registerValidators(validators ...validator.ValidatorInterface) {
	for _, validator := range validators {
		Validators[getType(validator)] = validator
	}
}

func init() {
	Validators = make(map[string]validator.ValidatorInterface)

	registerValidators(
		new(validator.CurrencyValidator),
		new(validator.TransactionTypeValidator),
	)
}

func validation(ctx *context.Context, method string, value interface{}) {
	if Validators[method].Validate(value) == false {
		resp := response.Response{
			Success: false,
			Data:    nil,
			Error:   "Error validate - " + method,
			Code:    400,
		}

		resp.ServeJSON(ctx)
	}
}
