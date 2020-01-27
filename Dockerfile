FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

RUN adduser -D -g '' scinna

WORKDIR $GOPATH/src/github.com/scinna/server/

COPY . .

RUN go mod download
RUN go mod verify

# @TODO ARM version
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hello

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/bin/scinna /go/bin/scinna

USER scinna

ENTRYPOINT ["/go/bin/scinna"]
