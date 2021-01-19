.PHONY: client client_watch server server_watch docker_build

clean:
	rm -f go-app.out server/web/app.wasm web/app.wasm

client:
	cd client && GOARCH=wasm GOOS=js go build -o ./../web/app.wasm
client_watch:
	cd client && GOARCH=wasm GOOS=js gowatch

server:
	cd server && go build -o ./../go-app.out
server_watch:
	cd server && gowatch

run:
	./go-app.out

docker_build:
	docker build -t go-app .
docker_run:
	docker run -it -p 8000:8000 go-app