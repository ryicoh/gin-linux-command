FROM golang:1.12-alpine3.9 as builder

WORKDIR /go/src
COPY main.go .

RUN apk update && \
    apk upgrade && \
    apk add build-base git && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go get -u github.com/gin-gonic/gin && \
    go build -gcflags "all=-N -l" -o /go/bin/school-linux

FROM alpine:3.9

RUN apk update && \
    apk add vim bash git

COPY --from=builder /go/bin/school-linux /usr/bin/school-linux

WORKDIR /root

EXPOSE 1313

CMD ["school-linux"]
