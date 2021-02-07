.PHONY: build-server

build-server: 
	go build -o ./dist/bpm-server ./cmd/server/...