[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:3ef15f5494cab4b8 -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# simplenHttpInterface_userMySql

三个各自独立、互不兼容的 Go HTTP 示例被打包在同一个包里：一个分页返回用户列表的 MySQL 接口、一个 hello-world 应答器，以及一个登录桩程序，三者从来没有作为同一个程序一起编译通过过。

[![License](https://img.shields.io/github/license/anyingiit/simplenHttpInterface_userMySql)](LICENSE)

[报告问题](https://github.com/anyingiit/simplenHttpInterface_userMySql/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/simplenHttpInterface_userMySql/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

仓库的 `main/` 目录下是三个互不相关的单文件 Go 程序，而不是一个整体应用，并且整个仓库里没有任何 `go.mod` 把它们绑定成一个模块。`main/main.go` 是仓库自己记录的唯一入口点（`facts.json` 里唯一的 entry point）：它在 8000 端口启动一个 HTTP 服务，提供 `/user_info` 接口，读取 `limit` 和 `page` 两个查询参数，对 MySQL 数据库中的 `user_info` 表做分页 `SELECT * FROM user_info` 查询，并把结果以 JSON 形式返回。`main/simInterface.go` 是第二个、与前者无关的程序：一个运行在 8080 端口的 hello-world HTTP 服务，唯一的路由 `/` 只会返回一句固定的问候语。`main/simplenInternetFace2.go` 是第三个：一个绑定在 `127.0.0.1:8080` 上的登录桩程序，把提交的用户名和密码与一个写死的测试账号比对，返回成功或失败的 JSON 状态码。

这三个文件都声明为 `package main`，各自还都定义了自己的 `func main`，所以它们不可能被一起编译：执行 `go build ./main/` 会因为 `main redeclared in this block` 而失败。`main/simplenInternetFace2.go` 即便单独编译也无法通过：它调用了 `NewBaseJsonBean`，而这个构造函数只存在于 `main/BaseJsonBean.go` 的注释里，该文件的全部内容都被注释掉了。只有 `main/main.go` 和 `main/simInterface.go` 各自单独编译时能够成功。

`main/main.go` 里的 `connSql` 函数还把数据库的用户名、密码和主机地址直接写死在源码中，而不是从配置或环境变量读取。这个凭据应当被视为已经泄露——从第二个提交起它就一直留在本仓库的历史记录里——请不要再复用它；应当轮换数据库密码，并把新密码改为从配置读取。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/simplenHttpInterface_userMySql/issues)。

## 开始使用

### 环境要求

- 一套 Go 工具链。仓库里任何地方都没有 `go.mod`，所以既不会锁定 Go 版本，也不会自动拉取依赖。
- 一个正在运行的 MySQL 服务器，其中要有 `user_info` 表（包含 `id`、`name`、`telephone`、`age` 这几列），并且要能被运行 `main/main.go` 的机器访问到——这是本仓库里唯一会访问数据库的部分。

### 安装

```sh
git clone https://github.com/anyingiit/simplenHttpInterface_userMySql.git
cd simplenHttpInterface_userMySql
go mod init simplenHttpInterface_userMySql
go get github.com/go-sql-driver/mysql
go build main/main.go
```

不存在整体模块级别的构建：执行 `go build ./main/` 会失败，因为 `main/simInterface.go` 和 `main/simplenInternetFace2.go` 也各自定义了自己的 `func main`。像上面这样单独构建 `main/main.go`，是唯一能得到可运行二进制文件的方式。

## 用法

```sh
./main &
curl "http://localhost:8000/user_info?limit=10&page=1"
```

这会启动上一步构建出的服务，并向它请求 `user_info` 表的第一页、每页 10 条数据。返回的 JSON 中带有 `code`、`message`，以及一个 `data` 对象，其内部的 `data` 字段才是真正的行数据列表。如果请求缺少 `limit` 或 `page`，服务会返回一句简短的拒绝提示而不是 JSON；并且无论哪种情况，每个请求都会先等待半秒钟才作答——这两点都写死在 `main/main.go` 里，无法配置。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/simplenHttpInterface_userMySql](https://github.com/anyingiit/simplenHttpInterface_userMySql)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
