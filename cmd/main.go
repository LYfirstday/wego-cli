package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/LYfirstday/wego-cli/internal/utils"
	"github.com/LYfirstday/wego-cli/internal/variables"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func ValidateLocalYamlPath() {
	cmdPath, _ := os.Getwd()
	localFilePath := strings.Join([]string{cmdPath, "wego.yaml"}, variables.FileUriSeparator)
	localFilePath2 := strings.Join([]string{cmdPath, "wego.yml"}, variables.FileUriSeparator)
	if isExist, _ := utils.IsPathExists(localFilePath); isExist {
		utils.ValidateConfiguration("./wego.yaml")
	} else if isExist, _ := utils.IsPathExists(localFilePath2); isExist {
		utils.ValidateConfiguration("./wego.yml")
	} else {
		fmt.Println("Need a wego.y[a]ml config file!")
		os.Exit(0)
	}
}

func main() {

	app := &cli.App{
		Name:    "wego",
		Usage:   "Frontend templates cli",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "Show pages、components & projects templates",
				Subcommands: []*cli.Command{
					{
						Name:    "page",
						Usage:   "wego s page      Show all page-templates",
						Aliases: []string{"page"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "config",
								Usage:   "Specify the profile",
								Aliases: []string{"c"},
							},
						},
						Action: func(cCtx *cli.Context) error {
							configFilePath := cCtx.Value("config")
							if configFilePath != "" {
								if str, ok := configFilePath.(string); ok {
									utils.ValidateConfiguration(str)
								} else {
									fmt.Println("Config param is not a valid value!")
									os.Exit(0)
								}
							} else {
								ValidateLocalYamlPath()
							}

							// 获取仓库wego.yaml文件配置内容
							configFile := utils.GetConfigFileInfo()
							answers := struct {
								PageTemplate string
								PageName     string
							}{}
							// 选择一个page模板，并下载到本地
							utils.GivenSurveySelect("PageTemplate", "Chose a page", configFile.GetComponentOrPageNames("page"), &answers)

							selectPageName := answers.PageTemplate

							var question = []*survey.Question{
								{
									Name:     "PageName",
									Prompt:   &survey.Input{Message: "What is your page name?"},
									Validate: survey.Required,
								},
							}
							inputErr := survey.Ask(question, &answers)

							if inputErr != nil {
								utils.LogError("Input page name error: ", inputErr)
							}
							// 命令行展示的name包含description了，这里需要将description去掉
							pName := strings.Split(selectPageName, " ")[0]
							utils.RunWriteJob("pages", pName, answers.PageName)
							configFile.LoadDependencies("page", pName)
							return nil
						},
					},
					{
						Name:    "component",
						Usage:   "wego s com      show all component-templates",
						Aliases: []string{"com"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "config",
								Usage:   "Specify the profile",
								Aliases: []string{"c"},
							},
						},
						Action: func(cCtx *cli.Context) error {
							configFilePath := cCtx.Value("config")
							if configFilePath != nil && configFilePath != "" {
								if str, ok := configFilePath.(string); ok {
									utils.ValidateConfiguration(str)
								} else {
									fmt.Println("Config param is not a valid value!")
									os.Exit(0)
								}
							} else {
								ValidateLocalYamlPath()
							}
							// 获取仓库wego.yaml文件配置内容
							configFile := utils.GetConfigFileInfo()

							answers := struct {
								ComTemplate string
							}{}
							// 选择一个组件模板，并下载到本地
							utils.GivenSurveySelect("ComTemplate", "Chose a component", configFile.GetComponentOrPageNames("component"), &answers)
							selectComName := answers.ComTemplate
							// 命令行展示的name包含description了，这里需要将description去掉
							comName := strings.Split(selectComName, " ")[0]
							utils.RunWriteJob("components", comName)
							configFile.LoadDependencies("component", comName)
							return nil
						},
					},
					{
						Name:    "project",
						Usage:   "wego s pro      Show all project-templates",
						Aliases: []string{"pro"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "config",
								Usage:   "Specify the profile",
								Aliases: []string{"c"},
							},
						},
						Action: func(cCtx *cli.Context) error {
							configFilePath := cCtx.Value("config")
							if configFilePath != nil && configFilePath != "" {
								if str, ok := configFilePath.(string); ok {
									utils.ValidateConfiguration(str)
								} else {
									fmt.Println("Config param is not a valid value!")
									os.Exit(0)
								}
							} else {
								ValidateLocalYamlPath()
							}
							// 获取仓库wego.yaml文件配置内容
							configFile := utils.GetConfigFileInfo()

							answers := struct {
								ProjectTemplate string
								ProName         string
							}{}

							// 选择一个工程模板，并下载到本地
							utils.GivenSurveySelect("ProjectTemplate", "Chose a project", configFile.GetProjectNames(), &answers)
							selectProName := answers.ProjectTemplate

							var question = []*survey.Question{
								{
									Name:     "ProName",
									Prompt:   &survey.Input{Message: "What is your project name?"},
									Validate: survey.Required,
								},
							}
							inputErr := survey.Ask(question, &answers)

							if inputErr != nil {
								utils.LogError("Input page name error: ", inputErr)
							}
							// 命令行展示的name包含description了，这里需要将description去掉
							utils.RunWriteJob("projects", strings.Split(selectProName, " ")[0], answers.ProName)
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		utils.LogError("App run error: ", err)
	}
}
