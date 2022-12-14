package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/LYfirstday/wego-cli/internal/constants"
	"github.com/LYfirstday/wego-cli/internal/variables"

	"gopkg.in/yaml.v3"
)

// GithubRepoInfo 抽象Github仓库，封装github仓库请求路径等功能
type GithubRepoInfo struct {
	Username    string `yaml:"username"`
	RepoName    string `yaml:"repoName"`
	GithubToken string `yaml:"githubToken"`
	IsPrivate   bool
}

var LocalConfigFile GithubRepoInfo

func ValidateConfiguration(configFilePath string) {
	configInstance := GithubRepoInfo{}
	finalPath := ""

	// absolute path
	if m, _ := regexp.MatchString(`^\/`, configFilePath); m {
		isExist, _ := IsPathExists(configFilePath)
		if isExist {
			finalPath = configFilePath
		} else {
			fmt.Println("No file named wego.y[a]ml! Path: ", configFilePath)
			os.Exit(0)
		}
	} else {
		//relative path
		cmdPath, cmdPathErr := os.Getwd()
		if cmdPathErr != nil {
			LogError("Get cmd path err: ", cmdPathErr)
		}
		fullPath := strings.Join([]string{cmdPath, strings.Split(configFilePath, "./")[1]}, variables.FileUriSeparator)
		isExist, _ := IsPathExists(fullPath)
		if isExist {
			finalPath = fullPath
		} else {
			fmt.Println("No file named wego.y[a]ml! Path: ", fullPath)
			os.Exit(0)
		}
	}

	content, readErr := os.ReadFile(finalPath)
	if readErr != nil {
		LogError("Read config file error: ", readErr)
	}
	parseYamlErr := yaml.Unmarshal(content, &configInstance)
	if parseYamlErr != nil {
		LogError("Parse config file error: ", parseYamlErr)
	}
	if configInstance.GithubToken != "" {
		configInstance.IsPrivate = true
	}
	LocalConfigFile = configInstance
}

func (repo *GithubRepoInfo) GetRepoYamlRequestUrl() string {
	requestUrl := []string{constants.API_PREFIX, repo.Username, repo.RepoName, "contents", "wego.yaml?ref=main"}
	return strings.Join(requestUrl, "/")
}
