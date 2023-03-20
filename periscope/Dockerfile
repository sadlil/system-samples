FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /go/samples/periscope
ADD . /go/samples

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/periscope ./service/...

# Application container
FROM alpine
COPY --from=builder /go/bin/periscope /go/bin/periscope
ENTRYPOINT ["/go/bin/periscope"]