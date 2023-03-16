all: build

build: clean
	./build.sh

clean:
	rm -rf build/linebot-go.linux.x64 build/linebot-go.windows.x64.exe build/linebot-go.mac.x64 build/linebot-go.mac.arm64

.PHONY: clean build all