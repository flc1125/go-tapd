# 🚀 Go-Tapd-SDK

![Supported Go Versions](https://img.shields.io/badge/Go-%3E%3D1.18-blue)
[![Package Version](https://badgen.net/github/release/flc1125/go-tapd/stable)](https://github.com/flc1125/go-tapd/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/flc1125/go-tapd)](https://pkg.go.dev/github.com/flc1125/go-tapd)
[![codecov](https://codecov.io/gh/flc1125/go-tapd/graph/badge.svg?token=QPTHZ5L9GT)](https://codecov.io/gh/flc1125/go-tapd)
[![Go Report Card](https://goreportcard.com/badge/github.com/flc1125/go-tapd)](https://goreportcard.com/report/github.com/flc1125/go-tapd)
[![lint](https://github.com/flc1125/go-tapd/actions/workflows/lint.yml/badge.svg)](https://github.com/flc1125/go-tapd/actions/workflows/lint.yml)
[![tests](https://github.com/flc1125/go-tapd/actions/workflows/test.yml/badge.svg)](https://github.com/flc1125/go-tapd/actions/workflows/test.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

The Go-Tapd-SDK is a Go client library for accessing the [Tapd API](https://www.tapd.cn/). 

It provides a simple way to interact with Tapd API.

## 📥 Installation

```bash
go get github.com/flc1125/go-tapd
```

## ✨ Features

see [features.md](features.md)

## 🔧 Usage

```go
package main

import (
	"context"
	"log"

	"github.com/flc1125/go-tapd"
)

func main() {
	client, err := tapd.NewClient("username", "password")
	if err != nil {
		log.Fatal(err)
	}

	// example: get labels
	labels, _, err := client.LabelService.GetLabels(context.Background(), &tapd.GetLabelsRequest{
		WorkspaceID: tapd.Ptr(123456),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("labels: %+v", labels)
}
```

## 📜 License

The MIT License (MIT). Please see [License File](LICENSE) for more information.