# Makefile

build: # build server
	go build -o ./.bin/app ./cmd/api/main.go

start: # start server
	./.bin/app

dev: # build and start server
	swag init --generalInfo ./cmd/api/main.go --output ./docs
	go build -o ./.bin/app ./cmd/api/main.go
	./.bin/app

docs:
	swag init --generalInfo ./cmd/api/main.go --output ./docs