APP_EXECUTABLE=mapp

build:
	GOARCH=amd64 GOOS=darwin go build -o ./out/${APP_EXECUTABLE} cmd/mapp/main.go

run: build
	./out/${APP_EXECUTABLE}

clean:
	go clean
	rm -rf ./out
