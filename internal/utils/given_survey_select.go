package utils

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func GivenSurveySelect(name string, msg string, options []string, answer interface{}) {
	if len(options) < 1 {
		fmt.Println("There is no any templates, please  make sure you have set github token if your repository is private!")
		os.Exit(0)
	}
	var question = []*survey.Question{
		{
			Name: name,
			Prompt: &survey.Select{
				Message: msg,
				Options: options,
			},
		},
	}
	selectErr := survey.Ask(question, answer)
	if selectErr != nil {
		LogError("Given survey select error: ", selectErr)
	}
}
