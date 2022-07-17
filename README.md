# go-lwapi

[![Build Status](https://travis-ci.org/gadost/go-lwapi.svg?branch=main)](https://travis-ci.org/gadost/go-lwapi) [![GoDoc](https://godoc.org/github.com/gadost/go-lwapi?status.svg)](https://godoc.org/github.com/gadost/go-lwapi)   [![codecov](https://codecov.io/gh/gadost/go-lwapi/branch/main/graph/badge.svg)](https://codecov.io/gh/gadost/go-lwapi)

## Install

```sh
go get -u github.com/gadost/go-lwapi@latest
```

## Example

```go
package main 

import (
    "github.com/gadost/go-lwapi"
)
func main() {
    api := lwapi.New("testtoken").DedicatedServers()
    api.BaseURL = server.URL

    resp, err := api.ServerReferenceUpdate(12345, &lwapi.Reference{Reference: "test-server"})
    ...
}
```
