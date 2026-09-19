<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# simplenHttpInterface_userMySql

Three independent, mutually incompatible Go HTTP demos bundled into one package: a paginated MySQL-backed user listing, a hello-world responder, and a login stub, none of which have ever built together as a single program.

**English** · [简体中文](README.zh-CN.md)

[![License](https://img.shields.io/github/license/anyingiit/simplenHttpInterface_userMySql)](LICENSE)

[Report a bug](https://github.com/anyingiit/simplenHttpInterface_userMySql/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/simplenHttpInterface_userMySql/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

The repository holds three unrelated, single-file Go programs under `main/`, not one application, and there is no `go.mod` anywhere to bind them into a module. `main/main.go` is the entry point the repository itself names (`facts.json`'s only recorded entry point): it starts an HTTP server on port 8000 and serves `/user_info`, which reads `limit` and `page` query parameters, runs a paginated `SELECT * FROM user_info` against a MySQL database, and returns the rows as JSON. `main/simInterface.go` is a second, unrelated program: a hello-world HTTP server on port 8080 whose only route, `/`, answers with a fixed greeting. `main/simplenInternetFace2.go` is a third: a login stub bound to `127.0.0.1:8080` that checks a submitted username and password against one hardcoded test account and returns a JSON success or failure code.

All three files declare `package main` and their own `func main`, so they cannot be built together — `go build ./main/` fails with `main redeclared in this block`. `main/simplenInternetFace2.go` does not build even on its own: it calls `NewBaseJsonBean`, a constructor that exists only inside a comment in `main/BaseJsonBean.go`, whose entire content is commented out. Only `main/main.go` and `main/simInterface.go` compile in isolation.

`main/main.go`'s `connSql` function also opens its MySQL connection with a username, password and host address written directly into the source, rather than read from configuration or an environment variable. Treat that credential as already exposed — it has been public in this repository's history since the second commit — and never reuse it; rotate the database password and read the next one from configuration instead.

See the [open issues](https://github.com/anyingiit/simplenHttpInterface_userMySql/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- A Go toolchain. The repository has no `go.mod` anywhere in the tree, so nothing pins a Go version or fetches a dependency automatically.
- A running MySQL server with a `user_info` table (`id`, `name`, `telephone`, `age` columns), reachable from wherever `main/main.go` runs — the only part of this repository that talks to a database.

### Installation

```sh
git clone https://github.com/anyingiit/simplenHttpInterface_userMySql.git
cd simplenHttpInterface_userMySql
go mod init simplenHttpInterface_userMySql
go get github.com/go-sql-driver/mysql
go build main/main.go
```

There is no module-wide build: `go build ./main/` fails, because `main/simInterface.go` and `main/simplenInternetFace2.go` each declare their own `func main` too. Building `main/main.go` on its own, as above, is the only path that produces a working binary.

## Usage

```sh
./main &
curl "http://localhost:8000/user_info?limit=10&page=1"
```

This starts the server built above and asks it for the first page of 10 rows from the `user_info` table. The JSON response carries `code`, `message`, and a `data` object whose own `data` field is the row list. A request missing `limit` or `page` gets a short refusal message instead of JSON, and every request waits half a second before it is answered — both are hardcoded in `main/main.go`, not configurable.

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/simplenHttpInterface_userMySql](https://github.com/anyingiit/simplenHttpInterface_userMySql)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
