FROM golang
COPY . /go/src/crypto-monitor
WORKDIR /go/src/crypto-monitor
RUN env GOOS=linux go build -ldflags="-s -w" -o bin/main main/main.go
# go get &&

FROM scratch
LABEL maintainer="Jordan Stewart <jordanstewart2428@gmail.com>"
COPY --from=0 /go/src/crypto-monitor/bin/main /app/bin/main
ENTRYPOINT [ "/app/bin/main" ]