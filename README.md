# Rocha

[![Build Status](https://travis-ci.com/vtfr/rocha.svg?branch=master)](https://travis-ci.com/vtfr/rocha)
[![codecov](https://codecov.io/gh/vtfr/rocha/branch/master/graph/badge.svg)](https://codecov.io/gh/vtfr/rocha)
[![Go Report Card](https://goreportcard.com/badge/github.com/vtfr/rocha)](https://goreportcard.com/report/github.com/vtfr/rocha)
[![GoDoc](https://godoc.org/github.com/vtfr/rocha?status.svg)](https://godoc.org/github.com/vtfr/rocha)

Rocha is Hyperledger Fabric Chaincode Router with Middleware capabilities

Currently in development

## Usage

```go
r := rocha.NewRouter().
    // Route method `QueryUsers` to QueryUsers function
    Handle("QueryUsers", QueryUsers)
    // Route method `SaveUser` receiving a string and integer parameters
    Handle("SaveUser", SaveUser,
        argsmw.Arguments(
            argsmw.String("name"),
            argsmw.Int("age"))).

func SaveUser(c rocha.Context) pb.Response {
    name := c.String("name")
    age := c.Int("age")

    if age < 18 {
        return shim.Error("can't register user")
    }

    stub := c.Stub()

    // ...
}
```
