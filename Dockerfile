FROM golang
COPY . /tmp
WORKDIR /tmp
RUN go get && env GOOS=linux go build -ldflags="-s -w" -o bin/main main/main.go

FROM scratch
LABEL maintainer="Jordan Stewart <jordanstewart2428@gmail.com>"
COPY --from=0 /tmp/bin/main /app/bin/main
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "./app/bin/main" ]