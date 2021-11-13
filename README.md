# Easegress Golang SDK

- [Easegress Golang SDK](#easegress-golang-sdk)
    - [Prerequisites](#prerequisites)
    - [Local Development](#local-development)
    - [Deploy and execute](#deploy-and-execute)

This is the [Golang](https://golang.org/) SDK for [Easegress](https://github.com/megaease/easegress), it can be used to extend the ability of Easegress.

## Prerequisites

The following assumes that a recent version of [Golang](https://golang.org/) and compiler [TinyGo](https://tinygo.org//) are installed.

## Local Development

1. Make a new directory and initialize a new module:

```bash
go mod init demo
```

2. Init a go file:

```bash
echo '
package main

import (
	"github.com/xmh19936688/easegress-go-sdk/easegress"
)

func init() {
	easegress.Log(easegress.Info, "hello sdk")
}

func main() {
	// keep empty
}
' > main.go
```

4. Install the dependency:

```bash
go mod tidy
```

5. Compile:

```bash
tinygo build -o demo.wasm -target wasi .
```

If everything is right, `demo.wasm` will be generated.

## Deploy and execute

Please refer [the documentation of `WasmHost`](https://github.com/megaease/easegress/blob/main/doc/wasmhost.md) for deploying and executing the compiled Wasm code.
