## Issue 6846

```
$ make up
// from another tab, $(docker-machine ip) on macOS or localhost on linux (platform without docker-machine).
$ curl $(docker-machine ip):10000/sample --verbose
*   Trying 192.168.99.100...
* TCP_NODELAY set
* Connected to 192.168.99.100 (192.168.99.100) port 10000 (#0)
> GET /sample HTTP/1.1
> Host: 192.168.99.100:10000
> User-Agent: curl/7.62.0
> Accept: */*
>
< HTTP/1.1 200 OK
< content-type: application/json
< x-envoy-upstream-service-time: 1
< grpc-status: 0
< grpc-message:
< content-length: 2
< date: Tue, 30 Jul 2019 23:09:25 GMT
< server: envoy
<
* Connection #0 to host 192.168.99.100 left intact
{}
```
