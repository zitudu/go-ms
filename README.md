# go-ms

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/zitudu/go-ms)
[![rcard](https://goreportcard.com/badge/github.com/json-iterator/go)](https://goreportcard.com/report/github.com/zitudu/go-ms)

Use this package to easily convert various time formats to milliseconds. Based on ms js <https://github.com/vercel/ms>

## Install

`go get github.com/zitudu/go-ms`

## Usage

```go
import "github.com/zitudu/go-ms"

age, err := ms.Parse("30days")
expiresIn := ms.MustParse("30days")
```
