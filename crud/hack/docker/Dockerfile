FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /go/samples/crud
ADD . /go/samples

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/crud-apiserver ./cmd/crud-apiserver/...

# Application container
FROM alpine
COPY --from=builder /go/bin/crud-apiserver /go/bin/crud-apiserver
ENTRYPOINT ["/go/bin/crud-apiserver"]