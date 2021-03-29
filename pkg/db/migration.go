package db

import (
	"wallet-test/models"
)

func migrate() {

	err := Msql.AutoMigrate(
		&models.User{},
		&models.Balance{},
		&models.Transaction{},
	)

	if err != nil {
		panic(err)
	}

}
