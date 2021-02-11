```
go run ./myserver.go
./traefik --configFile ./traefik.yml

# should be gzip
curl -vv whoami:8000/ --compressed

# should not be gzip
curl -vv whoami:8000/event --compressed
curl -vv whoami:8000/yolo --compressed
```