## gcm-proxy

gcm-proxy is GCM (Google Cloud Messaging) proxy written by golang.

## Useage


### Setup


```go
# gcm-proxy-server.go
package main

import (
	"log"

	gcm "github.com/Konboi/gcm-proxy"
)

func main() {
	config := &gcm.Config{
		Port:   6969,
		APIKey: "testapikey",
	}

	gcm_proxy, err := gcm.NewProxy(config)
	if err != nil {
		log.Fatalf("Initialize Error: %s", err.Error())
	}

	gcm_proxy.Run()
}

```


```
go run gcm-proxy-server.go
```

```
curl -d 'alert="test push"' -d 'token=123456' localhost:2222
```
