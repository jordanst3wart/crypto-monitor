.PHONY: build clean deploy scratch

build:
	env GOOS=linux /Users/jordanstewart/go/go1.21.5/bin/go build -ldflags="-s -w" -o bin/main main/functions.go main/main.go

clean:
	rm -rf ./bin

deploy: clean build
	scripts/deploy.sh

trial:
	go run main/functions.go main/main.go

build-container:
	docker build -t crypto-container .

list-container:
	docker image list

run-container:
	docker run crypto-container

test:
	scripts/test-build.sh

check-logs:
	scripts/check-logs.sh

