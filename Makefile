.PHONY: build clean deploy scratch

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

scratch:
	env GOOS=linux go build -ldflags="-s -w" -o bin/scratch scratch/main.go

trial:
	go run main/main.go

build-container:
	docker build -t crypto-container .

list-container:
	docker image list

run-container:
	docker run crypto-container

deploy:
	deploy.sh
