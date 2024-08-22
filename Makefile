run: build
	@./bin/app

build:
	@go build -o bin/app .

templ:
	templ generate --watch --proxy=http://127.0.0.1:4000 --open-browser=false

