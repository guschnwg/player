.PHONY: client client_watch server server_watch docker_build docker_run

client:
	GOARCH=wasm GOOS=js go build -o ./web/app.wasm cmd/client/main.go
client_watch:
	GOARCH=wasm GOOS=js gow build -v -o ./web/app.wasm cmd/client/main.go

server:
	go run cmd/server/main.go
server_watch:
	gow run cmd/server/main.go

docker_build:
	docker build -t go-app .
docker_run:
	docker run -p 8000:8000 -v ${shell pwd}/:/app --name go-app -it go-app