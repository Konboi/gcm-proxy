## gcm-proxy

gcm-proxy is GCM (Google Cloud Messaging) proxy written by golang.

## Useage


### Setup


```go
# gcm-proxy-server.go
impodrt(
    gcm_proxy "github.com/Konboi/gcm-proxy"
);


func main() {
    server := gcm_proxy.NewServer{
        Port: 22222,
        APIKey: 'abcdefg12345'
    };

    server.Run();
}
```


```
go run gcm-proxy-server.go
```

```
curl -d 'alert="test push"' -d 'token=123456' localhost:2222
```
