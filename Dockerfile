FROM golang:1.14 as builder

WORKDIR /go/src/nym-directory

COPY . .

RUN go build -ldflags "-s -w" -o nym-directory main.go

FROM ubuntu:18.04

COPY --from=builder /go/src/nym-directory/nym-directory /usr/bin/nym-directory

CMD ["nym-directory"]
