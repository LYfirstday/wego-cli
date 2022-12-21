# wego-cli

支持64位linux、windows、darwin操作系统
## wego-cli是什么

wego-cli是一个前端模板脚手架，能够支持你快速出创建、组织自己或团队的前端模板仓库，包括组件模板、页面模板和工程模板，提高日常开发效率；不局限于web端，安卓、IOS、pc、h5、小程序等多端组件、页面、工程模板都可以使用wego-cli组织、维护；wego-cli本身不提供任何模板。

## 为什么要有wego-cli

wego-cli和create-react-app、VUE cli等脚手架不同，要解决的问题完全不一样，react、VUE等前端框架脚手架面向的是全球用户，旨在快速构建出框架相关的app工程，脚手架非常复杂、精细; 而wego-cli要解决的是在企业内部开发中，如何更好的组织、维护企业内部相关技术栈的组件、页面、工程模板，并且能够快速将模板使用到日常开发中；企业内部开发通常是统一的技术栈和开发规范，在开发过程中使用统一规范的模板，能够有效降低维护成本，提高开发效率。

## 如何使用

### 第一步

在Github上创建一个模板仓库，公共的或私有的都可以(如果是私有仓库，需要在本地wego.yaml中配置githubToken，读权限)，模板仓库的目录遵循下面规范(目前不支持自定义模板文件夹名称)：

```text
  templates
    components
    pages
    projects
  wego.yaml
```

模板仓库本身也可以是一个前端工程，只要仓库根目录下有上述文件、文件夹。

templates: 存放组件、页面、工程模板的文件夹

components: 存放组件级模板

pages: 存放页面级模板

wego.yaml: 声明模板仓库详细信息和组件依赖关系(避免出现循环依赖)

一个wego.yaml模板信息声明依照下面规范:

```yaml
components:
  -
    name: ProductCategorySelector
    description: 商品分类选择器组件
    dependencies: []
  -
    name: ShipmentSettingModal
    description: 发货物流弹框组件
    dependencies: []
  -
    name: ProductBaseInfoForm
    description: 创建商品基本信息表单组件
    dependencies: [ProductCategorySelector]
  -
    name: MultiLevelComponents
    description: 多级依赖关系组件
    dependencies: [ProductBaseInfoForm]
pages:
  -
    name: product-list
    description: 商品列表页面
    dependencies: [ShipmentSettingModal]
  -
    name: create-product
    description: 创建商品表单页面
    dependencies: [ProductBaseInfoForm]
  -
    name: multi-level-dependencies-page
    description: 多级依赖关系页面
    dependencies: [MultiLevelComponents]
projects:
  -
    name: react-central-web-app
    description: react 管理端前端工程模板
  -
    name: next-h5-app
    description: next SSR h5前端工程模板
  -
    name: vue-app
    description: vue h5前端工程模板
```

模板文件夹下面的模板名称必须和wego.yaml中声明的信息一一对应，并且大小写敏感，每一个模板都是单独的文件夹，不要保存成单独的文件。

components、pages、projects分别对应组件、页面、工程模板文件夹，不支持自定义文件夹名称。

| Property    | Is required | Description                                                  |
| ----------- | :---------- | ------------------------------------------------------------ |
| name    | true        | 模板的名称，必须和templates中模板名称一一对应，并且大小写敏感                                                 |
| description    | false        | 对模板用途的解释说明                                                 |
| dependencies     | false        | 组件、页面模板所依赖的组件，依赖项只能是components中的组   件                                    |

以上是使用wego-cli的前置条件，必须有个模板仓库，查看 [示例模板仓库](https://github.com/LYfirstday/wego-cli-templates-example)

### 第二步

下载wego-cli

```shell
  npm install wego-cli [-g]
```

```shell
  yarn add wego-cli [global]
```

下载完成后，执行wego -v[-h]，如果出现以下结果证明安装成功

如果将wego-cli安装在本地npm依赖中，请执行:
```shell
npx wego -h[-v]
```

如果将wego-cli安装在全局npm依赖中，请执行:
```shell
wego -h[-v]
```
![Image text](https://raw.githubusercontent.com/LYfirstday/wego-cli/main/images/is_success.png)

### 第三步

本地创建一个空文件夹，并且创建wego.yaml文件，用于声明模板仓库信息，本地wego.yaml需要填写三条数据:

```yaml
username:
repoName:
githubToken:
```
| Property    | Is required | Description                                                  |
| ----------- | :---------- | ------------------------------------------------------------ |
| username    | true        | github账号名称                                      |
| repoName    | true        | 模板仓库名称              |
| githubToken     | false        | github token(读权限)，如果模板仓库为私有的则需要有读权限的authToken|

在wego.yaml文件路径下打开shell，执行wego-cli命令，wego-cli命令一览表：

| Command     | Description                                                  |
| ----------- | ------------------------------------------------------------ |
| wego -h     | 查看wego-cli 命令工具帮助信息                                    |
| wego -v     | 查看wego-cli 当前版本                                          |
| wego s[show] [options] | 列出模板仓库对应模板列表                              |

options 支持三种，对应模板仓库中的components、pages、projects

| Options     | Description                                                  |
| ----------- | ------------------------------------------------------------ |
| page[page]     | 列出'页面'模板                                    |
| com[components]     | 列出'组件'模板                                         |
| pro[project] | 列出'工程'模板                              |
| c[config] | 指定wego模板仓库配置文件，例如：wego s page config [xxx.yaml]                             |

例子:
1. 列出所有页面模板
```shell
wego s page
```
2. 列出所有组件模板
```shell
wego s com
```
3. 列出前端工程模板
```shell
wego s pro
```

执行命令后，使用箭头按键选择需要下载的模板，所有模板下载到本地的文件目录都是以当前执行shell命令的目录路径为基准，components中的模板会下载到: {{ cmdPath }}/src/components/ 目录下；pages中的模板会下载到: {{ cmdPath }}/src/pages/ 目录下；projects中的模板会下载到: {{ cmdPath }}/ 目录下；目前不支持自定义下载至目标文件路径；如果组件、页面模板文件夹在本地已存在则会略过(不覆盖原文件夹文件)。

![Image text](https://raw.githubusercontent.com/LYfirstday/wego-cli/main/images/show_all.png)

## Q & A

### 1. 什么样的组件模板适合用wego-cli组织、维护？

解决公司某个业务下某一类问题的组件(原则上所有模板组件都可以放这里)；

一般维护在wego中的组件，通常和公司业务场景有轻微耦合，这种组件解决公司特定业务下某个场景的需求，这种组件一般没有第三方可用的组件，都是自己或团队进行开发、实现，经过线上验证，后期优化，功能稳定后提炼、封装成组件，供其他团队或其他项目在相似的业务场景下使用。

### 2. 为什么不直接封装成组件库，发布到(私有)npm上维护？

一开始我也是这么做的，将业务组件发布到私有仓库，但是随着需求变更、UI更改，最初发布上去的业务组件已经没法用了，而且每次变更都要经历：修改 -> 测试 -> 发布 -> 升级组件版本的过程，效率极低而且麻烦；经常会出现，开发出组件A，但是某个场景下只用到组件A 70%的功能(核心功能不变)，如果为了这个场景专门修改组件A，支持新的场景，改的次数多了组件A就没法维护了，充满大量冗余的判断代码；不如把组件A稳定版源码下载下来做减法，直接修改组件源码、样式，达到快速支持业务场景的目的；你可能又要问为什么不用高阶组件将组件A包裹起来，高阶组件支持新场景，底层还是使用组件A，简单的组件可以这样做，但是对于稍微复杂点的业务组件，使用高阶组件到后面还是会面临难以维护的问题，冗余代码太多，尤其牵扯到组件样式定制、修改。

企业内部npm仓库适合发布比较通用的、底层的组件和工具类包，除非公司有专门团队维护npm组件包。

### 3. 为什么不好好设计下，把业务组件设计的扩展性更好、适用场景更多？

计划赶不上变化，骚年！

### 4. Error: API rate limit exceeded for xxxxx (But here's the good news: Authenticated requests get a higher rate limit. Check out the documentation for more details.)

这是因为github对没有验证的ip地址每小时有请求数量限制，详情请[查看](https://docs.github.com/zh/rest/rate-limit?apiVersion=2022-11-28)
解决方案：添加githubToken，也就无论你的模板仓库是共有或私有的都最好添加上有读权限的token
