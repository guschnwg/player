.PHONY: client client_watch server server_watch docker_build docker_run

client:
	GOARCH=wasm GOOS=js go build -ldflags "-s -w" -o ./web/app.wasm cmd/client/main.go
client_watch:
	GOARCH=wasm GOOS=js gow build -v -o ./web/app.wasm cmd/client/main.go

client_compress_brotli:
	brotli web/app.wasm
client_compress_gzip:
	gzip web/app.wasm --best --keep --force

server:
	go build -ldflags "-s -w" -o ./app.out cmd/server/main.go
server_watch:
	PORT=8000 gow run cmd/server/main.go

docker_build:
	docker build -t go-app .
docker_build_tinygo: # does not work, maybe one day
	docker run --rm -v $(shell pwd):/src -w /src tinygo/tinygo-dev tinygo build -o wasm.wasm -target=wasm cmd/client/main.go
docker_run:
	docker run -p 8000:8000 --env PORT=8000 --name go-app -it --rm go-app

heroku_push:
	heroku container:push web
heroku_deploy:
	heroku container:release web
heroku_logs:
	heroku logs --tail

vecty_build:
	GOARCH=wasm GOOS=js go build -ldflags "-s -w" -o ./gh-pages/vecty.wasm cmd/vecty/main.go

vugu_build:
	go generate cmd/vugu/generate.go

	GOARCH=wasm GOOS=js go build -ldflags "-s -w" -o ./gh-pages/vugu.wasm cmd/vugu/main_wasm.go cmd/vugu/*_vgen.go

gh_pages:
	git add -A
	git commit --allow-empty -m "Deploy"
	git subtree split --prefix gh-pages -b gh-pages
	git push -f origin gh-pages:gh-pages
	git branch -D gh-pages
	git push