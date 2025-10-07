APP_EXECUTABLE=mapp

build:
	GOARCH=amd64 GOOS=darwin go build -o ${APP_EXECUTABLE} cmd/mac-apps/main.go

run: build
	./${APP_EXECUTABLE}

clean:
	go clean
	rm ${APP_EXECUTABLE}
