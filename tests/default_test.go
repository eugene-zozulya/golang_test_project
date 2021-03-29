package test

import (
	"path/filepath"
	"runtime"
	"testing"
	"wallet-test/pkg/balance"
	"wallet-test/pkg/validator"
	_ "wallet-test/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestCurrencyValidator(t *testing.T) {
	validator := new(validator.CurrencyValidator)

	valid_values := []interface{}{"USD", "EUR"}

	for _, v := range valid_values {
		if !validator.Validate(v) {
			t.Error("Error currency validator")
		}
	}

	invalid_values := []interface{}{"usd", "eur", "Eur", "Usd"}

	for _, v := range invalid_values {
		if validator.Validate(v) {
			t.Error("Error currency validator")
		}
	}
}

func TestTransactionTypeValidator(t *testing.T) {
	validator := new(validator.TransactionTypeValidator)

	valid_values := []interface{}{"withdrowal", "deposit"}

	for _, v := range valid_values {
		if !validator.Validate(v) {
			t.Error("Error transaction type validator")
		}
	}

	invalid_values := []interface{}{"Withdrowal", "Deposit", "WITHDROWAL", "DEPOSIT"}

	for _, v := range invalid_values {
		if validator.Validate(v) {
			t.Error("Error transaction type validator")
		}
	}
}

func TestBalanceWithdrowal(t *testing.T) {
	balance := new(balance.Balance)
	balance.Init(1, "EUR")
	balance.BeforeValue = 1.000
	balance.AfterValue = 1.000

	err := balance.Withdrowal(2.000)
	if err == nil {
		t.Error("Must be error. Insufficient funds")
	}

	if balance.AfterValue != balance.BeforeValue {
		t.Error("Invalid balance")
	}

	err = balance.Withdrowal(-1.000)
	if err == nil {
		t.Error("Must be error. Negative value")
	}

	if balance.AfterValue != balance.BeforeValue {
		t.Error("Invalid balance")
	}

	err = balance.Withdrowal(0.005)
	if err != nil {
		t.Error("Error found")
	}

	if balance.AfterValue != 0.995 {
		t.Error("Invalid balance")
	}
}

func TestBalanceDepositl(t *testing.T) {
	balance := new(balance.Balance)
	balance.Init(1, "EUR")
	balance.BeforeValue = 1.000
	balance.AfterValue = 1.000

	err := balance.Deposit(-1.000)
	if err == nil {
		t.Error("Must be error. Negative value")
	}

	if balance.AfterValue != balance.BeforeValue {
		t.Error("Invalid balance")
	}

	err = balance.Deposit(2.000)
	if err != nil {
		t.Error("Error found")
	}

	if balance.AfterValue == balance.BeforeValue {
		t.Error("Invalid balance")
	}
}
