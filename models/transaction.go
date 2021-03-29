package models

import (
	"database/sql/driver"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	parsed, err := time.Parse("02-Jan-06 15:04:05", string(s))
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	s := t.Format("02-Jan-06 15:04:05")
	return []byte(s), nil
}

func (j *Time) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		j.Time = t
	}
	return nil
}

func (j Time) Value() (driver.Value, error) {
	return j.Time, nil
}

type Transaction struct {
	gorm.Model

	UserID        uint    `json:"user_id"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	TimePlaced    Time    `json:"time_placed"`
	Type          string  `json:"type"`
	BalanceBefore float64
	BalanceAfter  float64
}
