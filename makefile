config=config.toml

all: build

build:
	./build.sh

clean:
	rm -rf build/linebot-go build/linebot-go.linux.x64 build/linebot-go.windows.x64.exe build/linebot-go.mac.x64 build/linebot-go.mac.arm64

run:
	docker-compose up --build | ./build/linebot-go --config $(config) http --port 8080 | ./ngrok http 8080

run-local:
	docker-compose up --build | go run main.go --config $(config) http --port 8080 | ./ngrok http 8080

.PHONY: clean build all