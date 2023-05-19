flags="-X main.version=0.1.0  -X 'linebot-go/cmd.goversion=$(go version)' -X linebot-go/cmd.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X linebot-go/cmd.githash=`git describe --always --long --dirty --abbrev=14`"

echo "$flags"

# # build current
go build -ldflags "$flags" -o build/linebot-go -o build/linebot-go

# build for Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$flags" -o build/linebot-go.linux.x64 main.go 

# build for Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$flags" -o build/linebot-go.windows.x64.exe main.go 

# build for Mac (intel)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$flags" -o build/linebot-go.mac.x64 main.go  

# build for Mac (apple)
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$flags" -o build/linebot-go.mac.arm64 main.go