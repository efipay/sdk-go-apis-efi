# gn-api-sdk-go

> A go library for integration of your backend with the payment services
provided by [Efipay](http://gerencianet.com.br).

[![Build Status](https://travis-ci.org/gerencianet/gn-api-sdk-go.svg)](https://travis-ci.org/gerencianet/gn-api-sdk-go)
[![Coverage Status](https://coveralls.io/repos/github/gerencianet/gn-api-sdk-go/badge.svg?branch=master)](https://coveralls.io/github/gerencianet/gn-api-sdk-go?branch=master)

## Installation

Install with:

```bash

$ go get github.com/efipay/sdk-go-apis-efi/src/efipay
$ go mod init github.com/efipay/sdk-go-apis-efi
```
## Tested with
```
go 1.8, 1.11.4, 1.16.5 and 1.19.2

```
## Basic usage

```go

import (
    "github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

credentials := map[string]interface{} {
    "client_id": "client_id",
    "client_secret": "client_secret",
    "sandbox": false,
    "timeout": 10,
}

efi := efipay.NewEfiPay(credentials)

body = {
    "items": [{
        "name": "Product 1",
        "value": 1000,
        "amount": 2,
    }],
    "shippings": [{
        "name": "Default Shipping Cost",
        "value": 100,
    }]
}

res, err := efi.CreateCharge(body)

```

## Examples

You can run the examples inside `examples` with
`$ go run example.go`:

```bash
$ go run charges/billet/createCharge.go
```

Just remember to set the correct credentials inside `examples/configs.go` before running.

## Tests

To run the tests, just run:

```bash
$ go test
```

## Changelog

[CHANGELOG](CHANGELOG.md)


## License

The library is available as open source under the terms of the [MIT License](LICENSE).
