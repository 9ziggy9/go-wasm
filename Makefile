GOOS=js
GOARCH=wasm

all: main.go
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o main.wasm main.go

serve: serve.go
	go run serve.go

clean:
	rm -rf app
