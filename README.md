# go-ms

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/zitudu/go-ms)
[![rcard](https://goreportcard.com/badge/github.com/json-iterator/go)](https://goreportcard.com/report/github.com/zitudu/go-ms)

Use this package to easily convert various time formats to milliseconds. This is a Go port of [the Javascript ms library](https://github.com/vercel/ms).

## Install

Requires Go 1.11 and above. Run this command

    go get github.com/zitudu/go-ms

## Usage

```go
import "github.com/zitudu/go-ms"

age, err := ms.Parse("30days") // 2592000000.000000
expiresIn := ms.MustParse("20 hrs") // 72000000.000000
ms.FormatShort(60000) // 1m
ms.FormatLong(-3 * 60000) // -3 minutes
```

## License

MIT
