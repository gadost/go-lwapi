# go-lwapi

[![Build Status](https://travis-ci.org/gadost/go-lwapi.svg?branch=main)](https://travis-ci.org/gadost/go-lwapi) [![GoDoc](https://godoc.org/github.com/gadost/go-lwapi?status.svg)](https://godoc.org/github.com/gadost/go-lwapi)   [![codecov](https://codecov.io/gh/gadost/go-lwapi/branch/main/graph/badge.svg)](https://codecov.io/gh/gadost/go-lwapi) [![Go Report Card](https://goreportcard.com/badge/github.com/gadost/go-lwapi)](https://goreportcard.com/report/github.com/gadost/go-lwapi)

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
    // Dedicated Server AAPI
    api := lwapi.New("testtoken").DedicatedServers()

    resp, err := api.ServerReferenceUpdate(12345, &lwapi.Reference{Reference: "test-server"})
    ...

    // Virtual Server API
    api := lwapi.New("testtoken").VirtualServers()

    resp, err := api.VirtualServer(12345)
    ...
}
```

## Query params

```go
api := lwapi.New("testtoken").DedicatedServers()

q := make(map[string]interface{})
q["limit"] = 20
q["privateNetworkEnabled"] = "true"
q["ip"] = "127.0.0.1"

resp, err := api.Servers(q)
...
```

## Time helper ISO-8601

```go
 to := FormatISO8601(time.Now().Add(60*time.Day))
```

returns time in ISO-8601 format
