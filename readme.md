# Pez

A simple gRPC application that mimics a key-value store.

This project is implemented using [gRPC](https://grpc.io/) and [Cobra](https://pkg.go.dev/github.com/spf13/cobra).

## Prerequisites

In order to build this program you need [Go](https://go.dev/doc/install) installed. You can compile the program and produce the `pez` binary by executing:

```
$ go build
```

## Usage

Shell 1:

```
$ pez start
```

Shell 2:

```
$ pez put foo bar
2022/01/12 17:12:43 Assigned (foo, bar)

$ pez get foo
2022/01/12 17:12:46 Retrieved (foo, bar)
```

## Commands

The program commands that can be used:

| Command     | Description                                                      | Example                |
| :---------- | :----------------------------------------------------------------| :----------------------|
| `start`     | Starts the pez server                                            | `pez start`            |
| `put`       | Store a key, value pair                                          | `pez put foo bar`      |
| `get`       | Retrieve a key, value pair                                       | `pez get foo`          |

## Flags

The program flags that can be used:

| Flag                  | Description                                                      | Example            |
| :---------------------| :----------------------------------------------------------------| :----------------- |
| `-p`, `--port`        | Specifies which port to connect to                               | `-p 8080`          |

## Compiling the Protocol Buffer records

With [Protoc](https://grpc.io/docs/protoc-installation/) installed, execute:

```
$ protoc api/v1/*.proto \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--proto_path=.
```

Alternatively you can execute `make compile` in the root directory of this project if you have [Make](https://www.gnu.org/software/make/) installed.
