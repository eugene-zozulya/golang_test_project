package balance

import (
	"errors"
	"math/big"
	"wallet-test/models"
	"wallet-test/pkg/db"
)

type Balance struct {
	BeforeValue float64
	AfterValue  float64
	UserID      uint
	Currency    string
}

func (b *Balance) Get() bool {

	var balance models.Balance

	result := db.Msql.Where("user_id = ?", b.UserID).Where("currency = ?", b.Currency).First(&balance)

	if (result.RowsAffected == 0) || (result.Error != nil) {
		return false
	}

	b.BeforeValue = balance.Value
	b.AfterValue = balance.Value

	return true
}

func (b *Balance) Withdrowal(value float64) error {
	if value < 0 {
		return errors.New("The value can not be negative")
	}

	before := (&big.Float{}).SetFloat64(b.BeforeValue)

	withdrowal_value := (&big.Float{}).SetFloat64(value)

	result := new(big.Float).Sub(before, withdrowal_value)

	val, _ := result.Float64()

	if val < 0 {
		return errors.New("Insufficient funds")
	}

	b.AfterValue = val

	return nil
}

func (b *Balance) Deposit(value float64) error {
	if value < 0 {
		return errors.New("The value can not be negative")
	}

	before := (&big.Float{}).SetFloat64(b.BeforeValue)

	add_value := (&big.Float{}).SetFloat64(value)
	result := new(big.Float).Add(before, add_value)

	b.AfterValue, _ = result.Float64()

	return nil
}

func (b *Balance) Init(user_id uint, currency string) {
	b.Currency = currency
	b.UserID = user_id
}

func (b *Balance) Save(rec models.Transaction) error {

	tx := db.Msql.Begin()

	result := tx.Model(models.Balance{}).Where("user_id = ?", b.UserID).Where("currency = ?", b.Currency).Update("value", b.AfterValue)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	rec.BalanceBefore = b.BeforeValue
	rec.BalanceAfter = b.AfterValue
	result = tx.Create(&rec)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
