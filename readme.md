# Test websocket

## Test on docker
1. Start kong
```shell
docker compose up -d
```
2. Start ws server
```shell
cd ws && go run server/main.go
```
3. Start ws server
```shell
cd ws && go run client/main.go
```