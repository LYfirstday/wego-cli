package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/LYfirstday/wego-cli/internal/constants"
	"github.com/LYfirstday/wego-cli/internal/variables"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/color"
)

func CreateFolder(createPath string, item GithubResponseFile) {
	os.MkdirAll(createPath, os.ModePerm)
	client := resty.New()
	request := client.R().ForceContentType(constants.JSON_CONTENT_TYPE)
	if LocalConfigFile.IsPrivate {
		request.SetAuthToken(LocalConfigFile.GithubToken)
	}
	response, err := request.Get(item.Url)
	if err != nil {
		LogError("Request folder error: ", err)
	}
	fileRes := []GithubResponseFile{}
	parseJsonErr := json.Unmarshal(response.Body(), &fileRes)
	if parseJsonErr != nil {
		LogError("Parse json error: ", parseJsonErr)
	}
	WriteToLocal(fileRes, createPath)
	defer func() {
		color.Println(color.Green(createPath), ": create done!")
		variables.WriteWg.Done()
	}()
}

func CreateFile(createPath string, item GithubResponseFile) {
	client := resty.New()
	request := client.R().ForceContentType(constants.JSON_CONTENT_TYPE)
	if LocalConfigFile.IsPrivate {
		request.SetAuthToken(LocalConfigFile.GithubToken)
	}
	response, resErr := request.Get(item.Url)

	if resErr != nil {
		LogError("Request file error: ", resErr)
	}
	fileRes := GithubResponseFile{}
	parseJsonErr := json.Unmarshal(response.Body(), &fileRes)
	if parseJsonErr != nil {
		LogError("Parse config file error: ", parseJsonErr)
	}

	filePath := strings.Join([]string{createPath, item.Name}, variables.FileUriSeparator)
	file, err := os.Create(filePath)

	if err != nil {
		LogError("Create file error: ", err)
	}

	decStr, _ := base64.StdEncoding.DecodeString(fileRes.Content)
	_, wErr := file.WriteString(string(decStr))
	if wErr != nil {
		LogError("Write file error: ", wErr)
	}
	defer func() {
		color.Println(color.Green(filePath), ": write done!")
		variables.WriteWg.Done()
		file.Close()
	}()
}

func WriteToLocal(fileList []GithubResponseFile, rp string) {
	os.MkdirAll(rp, os.ModePerm)
	for _, item := range fileList {
		if item.Type == "dir" {
			variables.WriteWg.Add(1)

			go CreateFolder(strings.Join([]string{rp, item.Name}, variables.FileUriSeparator), item)
		}
		if item.Type == "file" {
			variables.WriteWg.Add(1)

			go CreateFile(rp, item)
		}
	}
}

func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// loadType取值: pages, components, projects
func RunWriteJob(loadType string, folderName string, newName ...string) {
	timeBegin := time.Now().UnixNano() / 1e6
	localFolderName := folderName
	// 获取当前命令执行路径，根据模板类型拼接出目标文件写入路径，如果是components类型，需要判断目标文件夹下是否存在这个组件
	// 如果存在，打印一条warning提醒
	targetPath, _ := os.Getwd()
	var rootPath string
	// 如果是创建page或project，会让开发输入自定义文件名，创建文件夹时用新的名称
	if len(newName) > 0 {
		localFolderName = newName[0]
	}
	// 如果创建project则在执行命令的文件中创建，如果创建page或component，默认创建到src下面的pages或components文件夹中
	if loadType == "projects" {
		rootPath = strings.Join([]string{targetPath, localFolderName}, variables.FileUriSeparator)
	} else {
		rootPath = strings.Join([]string{targetPath, "src", loadType, localFolderName}, variables.FileUriSeparator)
	}

	switch {
	case loadType == "components":
		isExist, _ := IsPathExists(rootPath)

		if isExist {
			warningMsg := []string{"Warning: ", folderName, " ", "already exist!!!"}
			color.Println(color.Red(strings.Join(warningMsg, "")))
			break
		}
		fallthrough
	default:
		// 根据名称，拉取远程仓库模板数据
		client := resty.New()
		// 模板请求地址
		url := strings.Join([]string{constants.API_PREFIX, LocalConfigFile.Username, LocalConfigFile.RepoName, "contents/templates", loadType, folderName}, "/")
		request := client.R().ForceContentType(constants.JSON_CONTENT_TYPE)
		if LocalConfigFile.IsPrivate {
			request.SetAuthToken(LocalConfigFile.GithubToken)
		}
		// 返回组件文件夹数据
		response, err := request.Get(url)

		if err != nil {
			fmt.Printf("Request %s template error: %s", loadType, err)
			os.Exit(0)
		}

		// 将模板文件夹数据解析到fileRes中，其实就是模板文件夹下的所有文件集合
		fileRes := []GithubResponseFile{}
		parseJsonErr := json.Unmarshal(response.Body(), &fileRes)
		if parseJsonErr != nil {
			fmt.Printf("Parse %s template json error: %s", loadType, parseJsonErr)
			os.Exit(0)
		}

		os.MkdirAll(rootPath, os.ModePerm)
		WriteToLocal(fileRes, rootPath)
		variables.WriteWg.Wait()
		costTime := (time.Now().UnixNano() / 1e6) - timeBegin
		color.Println(color.Green("Done!"), color.Green(costTime), "ms")
	}
}
