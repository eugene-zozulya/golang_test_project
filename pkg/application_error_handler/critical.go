package application_error_handler

import (
	"fmt"
	"os"
)

func Error_handler() {
	if r := recover(); r != nil {
		fmt.Println("Critical error application. ", r)
		os.Exit(0)
	}
}
