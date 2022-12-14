package utils

import (
	"fmt"
	"os"
)

type callback func()

func LogError(msg string, err error, cb ...callback) {
	fmt.Println(msg, err)
	for _, fn := range cb {
		fn()
	}
	os.Exit(0)
}
