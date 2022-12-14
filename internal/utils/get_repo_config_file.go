// @Title utils
// @Description 工具包，放置Github仓库文件、yaml配置
// 文件等工具函数
package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/LYfirstday/wego-cli/internal/constants"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/color"
	"gopkg.in/yaml.v3"
)

type ComponentsAndPages struct {
	Name         string   `yaml:"name"`
	Description  string   `yaml:"description"`
	Dependencies []string `yaml:"dependencies"`
}
type Projects struct {
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
}

type RepoConfigFile struct {
	Components []ComponentsAndPages `yaml:"components"`
	Pages      []ComponentsAndPages `yaml:"pages"`
	Projects   []Projects           `yaml:"projects"`
}

func GetConfigFileInfo() RepoConfigFile {
	// 请求仓库wego配置文件
	client := resty.New()
	request := client.R().ForceContentType(constants.JSON_CONTENT_TYPE)
	if LocalConfigFile.IsPrivate {
		request.SetAuthToken(LocalConfigFile.GithubToken)
	}
	response, err := request.Get(LocalConfigFile.GetRepoYamlRequestUrl())
	if err != nil {
		LogError("Fetch config file error: ", err)
	}
	// 解析json res
	fileRes := GithubResponseFile{}
	parseJsonErr := json.Unmarshal(response.Body(), &fileRes)
	if parseJsonErr != nil {
		LogError("Parse config file error: ", parseJsonErr)
	}
	// 将加密文件内容解析出来并转为struct对象
	yamlFile := RepoConfigFile{}
	decStr, _ := base64.StdEncoding.DecodeString(fileRes.Content)
	parseYamlErr := yaml.Unmarshal(decStr, &yamlFile)
	if parseYamlErr != nil {
		LogError("Parse config yaml file error: ", parseYamlErr)
	}

	return yamlFile
}

// 从yaml文件中获取模板列表
// templateType取值：page or component
func (f *RepoConfigFile) GetComponentOrPageNames(templateType string) []string {
	var list []ComponentsAndPages
	if templateType == "page" {
		list = f.Pages
	} else {
		list = f.Components
	}

	if len(list) > 0 {
		components := make([]string, 0)
		for _, item := range list {
			info := item.Name
			if item.Description != "" {
				info = strings.Join([]string{item.Name, "    --    ", item.Description}, "")
			}
			if len(item.Dependencies) > 0 {
				depsStr := ""
				depsStr = strings.Join(item.Dependencies, ", ")
				depsStr = "[" + depsStr + "]"
				info = strings.Join([]string{info, "   deps: ", depsStr}, "")
			}
			components = append(components, info)
		}
		return components
	}

	return make([]string, 0)
}

// 从yaml文件中获取工程模板列表
func (f *RepoConfigFile) GetProjectNames() []string {
	if len(f.Projects) > 0 {
		pros := make([]string, 0)
		for _, pro := range f.Projects {
			pros = append(pros, strings.Join([]string{pro.Name, "    --     ", pro.Description}, ""))
		}
		return pros
	}
	return make([]string, 0)
}

func (f *RepoConfigFile) LoadDependencies(templateType string, templateName string) {
	var thisTemplateType ComponentsAndPages
	if templateType == "page" {
		for _, page := range f.Pages {
			if page.Name == templateName {
				thisTemplateType = page
				break
			}
		}
	} else {
		for _, com := range f.Components {
			if com.Name == templateName {
				thisTemplateType = com
				break
			}
		}
	}

	deps := thisTemplateType.Dependencies
	if len(deps) > 0 {
		fmt.Println(color.Red("Loading dependencies..."))
		for _, dep := range deps {
			RunWriteJob("components", strings.Split(dep, " ")[0])
		}
	}
}
