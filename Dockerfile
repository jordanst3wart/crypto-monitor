# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
FROM golang:alpine as builder
WORKDIR /go/src/crypto-monitor
COPY . .
LABEL maintainer="Jordan Stewart <jordanstewart2428@gmail.com>"
# Install git.
# Git is required for fetching the dependencies.
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
# Create appuser
RUN apk update && apk add --no-cache git && \
    go get -d -v ./... && \
    env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /go/bin/main main/main.go && \
    adduser -D -g '' appuser

#FROM scratch
FROM golang
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/main /go/bin/main
# USER appuser
#CMD [ "ls && /go/bin/main" ]
#CMD [ "sleep infinity" ]
CMD pwd && ls /go/bin/main
#CMD /go/bin/main