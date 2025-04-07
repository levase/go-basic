.SILENT:

# Сборка под windows
build-win:
	@set GOOS=windows
	@set GOARCH=amd64
	@go build -o bin\JSONBin.exe main.go

# Сборка под macOS (Intel x86_64)
build-mac:
	@set GOOS=darwin
	@set GOARCH=amd64
	@go build -o bin/JSONBin-mac-amd64 main.go

# Сборка под macOS (Apple arm64)
build-mac-arm:
	@set GOOS=darwin
	@set GOARCH=arm64
	@go build -o bin/JSONBin-mac-arm64 main.go