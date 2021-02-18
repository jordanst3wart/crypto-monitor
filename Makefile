.PHONY: build clean deploy scratch

build:
	# env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	# env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main/functions.go main/main.go

clean:
	rm -rf ./bin

deploy: clean build
	deploy.sh

trial:
	go run main/functions.go main/main.go

build-container:
	docker build -t crypto-container .

list-container:
	docker image list

run-container:
	docker run crypto-container

test:
	env GOOS=linux go test -ldflags="-s -w"
