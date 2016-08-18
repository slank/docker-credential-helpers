# THIS IS BROKEN

Just stashing it here in case I decide to pick it back up later

## Introduction

docker-credential-helpers is a suite of programs to use altenate sources of Docker credentials.

## Installation

Go to the [Releases](https://github.com/slank/docker-credential-helpers/releases) page and download the binary that works better for you. Put that binary in your `$PATH`, so Docker can find it.

### Building from scratch

The programs in this repository are written with the Go programming language. These instructions assume that you have previous knowledge about the language and you have it installed in your machine.

1 - Download the source and put it in your `$GOPATH` with `go get`.

```
$ go get github.com/docker/docker-credential-helpers
```

2 - Use `make` to build the program you want. That will leave any executable in the `bin` directory inside the repository.

```
$ cd $GOPATH/slank/docker-credentials-helpers
$ make subhelper
```

3 - Put that binary in your `$PATH`, so Docker can find it.

## Usage

Set the `credsStore` option in your `.docker/config.json` file with the suffix of the program you want to use. For instance, set it to `subhelper` if you want to use `docker-credential-subhelper`.

```json
{
  "credsStore": "subhelper"
}
```

### Available programs

1. subhelper: Provides a helper that can select other helpers based on the registry endpoint.

## License

MIT. See [LICENSE](LICENSE) for more information.
