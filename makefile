APP_EXECUTABLE = mapp

build:
	GOARCH=arm64 GOOS=darwin go build -o ./out/${APP_EXECUTABLE}-arm cmd/mapp/main.go
	GOARCH=amd64 GOOS=darwin go build -o ./out/${APP_EXECUTABLE}-intel cmd/mapp/main.go

run: build
	@ARCH=$$(uname -m); \
	if [ "$$ARCH" = "arm64" ]; then \
		./out/${APP_EXECUTABLE}-arm; \
	else \
		./out/${APP_EXECUTABLE}-intel; \
	fi

clean:
	go clean
	rm -rf ./out
