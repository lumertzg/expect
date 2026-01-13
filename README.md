# expect

[![Build Status](https://github.com/lumertzg/expect/actions/workflows/ci.yml/badge.svg)](https://github.com/lumertzg/expect/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/lumertzg/expect.svg)](https://pkg.go.dev/github.com/lumertzg/expect) [![Go Report Card](https://goreportcard.com/badge/github.com/lumertzg/expect)](https://goreportcard.com/report/github.com/lumertzg/expect)

An assertion library to simplify writing tests in Go.

## Features

- Zero dependencies
- Generic functions for type-safe comparisons
- Support for slices and maps
- Works with `*testing.T` and `*testing.B`
- Clear failure messages

## Installation

```bash
go get github.com/lumertzg/expect
```

## Usage

```go
package mypackage

import (
    "testing"

    "github.com/lumertzg/expect"
)

func TestExample(t *testing.T) {
    expect.Equal(t, 42, Calculate())
    expect.True(t, IsValid())
    expect.NoError(t, DoSomething())
}
```

See the [examples](./examples) folder for more usage examples.
