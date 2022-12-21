package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type callback func()

func LogError(msg string, err error, cb ...callback) {
	fmt.Println(msg, err)
	for _, fn := range cb {
		fn()
	}
	os.Exit(0)
}

type ResponseMsg struct {
	Message string `json:"message"`
}

func CheckResponse(response *resty.Response) {
	if response.StatusCode() != 200 {
		res := ResponseMsg{}
		parseJsonErr := json.Unmarshal(response.Body(), &res)
		if parseJsonErr != nil {
			LogError("Parse config file error: ", parseJsonErr)
		}
		fmt.Println("Error: ", res.Message)
		os.Exit(0)
	}
}
