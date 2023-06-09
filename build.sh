# Windows
GOOS=windows GOARCH=386 go build -o bin/conventional-windows-386.exe ./...
GOOS=windows GOARCH=amd64 go build -o bin/conventional-windows-amd64.exe ./...
GOOS=windows GOARCH=arm go build -o bin/conventional-windows-arm.exe ./...
GOOS=windows GOARCH=arm64 go build -o bin/conventional-windows-arm64.exe ./...

# MacOS
GOOS=darwin GOARCH=amd64 go build -o bin/conventional-mac-amd64 ./...
GOOS=darwin GOARCH=arm64 go build -o bin/conventional-mac-arm64 ./...

# Linux
GOOS=linux GOARCH=amd64 go build -o bin/conventional-linux-amd64 ./...
GOOS=linux GOARCH=arm go build -o bin/conventional-linux-arm ./...
