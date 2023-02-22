GOOS=js
GOARCH=wasm
BUILD_DIR=bin

all: main.go
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(BUILD_DIR)/main.wasm main.go

serve: serve.go
	go run serve.go

clean:
	rm -rf bin/main.wasm
