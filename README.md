# go-lwapi

```sh
go get -u github.com/gadost/go-lwapi@latest
```

## Example

```go
import (
    "github.com/gadost/go-lwapi"
)

api := lwapi.New("testtoken").DedicatedServers()
api.BaseURL = server.URL

resp, err := api.ServerReferenceUpdate(12345, &lwapi.Reference{Reference: "test-server"})
...
```
