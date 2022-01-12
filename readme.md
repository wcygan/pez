# Pez

A simple key-value store.

This project is implemented using [gRPC](https://grpc.io/) and [Cobra](https://pkg.go.dev/github.com/spf13/cobra).

## Prerequisites

In order to build this program you need [Go](https://go.dev/doc/install),
[Protoc](https://grpc.io/docs/protoc-installation/), and [Make](https://www.gnu.org/software/make/manual/make.html)
installed.

Once they are installed, you can compile the program and produce the `pez` binary by executing:

```
$ make compile
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
