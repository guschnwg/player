.PHONY: client client_watch server server_watch docker_build docker_run

client:
	GOARCH=wasm GOOS=js go build -ldflags "-s -w" -o ./web/app.wasm cmd/client/main.go
client_watch:
	GOARCH=wasm GOOS=js gow build -v -o ./web/app.wasm cmd/client/main.go

server:
	go build -ldflags "-s -w" -o ./app.out cmd/server/main.go
server_watch:
	PORT=8000 gow run cmd/server/main.go

run:
	PORT=8000 ./app.out

docker_build:
	docker build -t go-app .
docker_run:
	docker run -p 8000:8000 --name go-app -it -e PORT=8000 --rm go-app

heroku_push:
	heroku container:push web
heroku_deploy:
	heroku container:release web