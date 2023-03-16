# build for Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/linebot-go.linux.x64 main.go 

# build for Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/linebot-go.windows.x64.exe main.go 

# build for Mac (intel)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/linebot-go.mac.x64 main.go  

# build for Mac (apple)
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/linebot-go.mac.arm64 main.go 