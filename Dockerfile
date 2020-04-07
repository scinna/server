FROM golang
WORKDIR /app
EXPOSE 40000 1541 1635

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN go get github.com/go-delve/delve/cmd/dlv

# Temporarly? commented: something is wrong with go-nanoid that prevents it from being downloaded
# Since I push my
# "go: github.com/matoous/go-nanoid@1.1.0: parsing go.mod: go.mod1: usage: go 1.23"
RUN GO111MODULE=on go mod download 

CMD [ "dlv", "debug", "github.com/scinna/server", "--listen=:40000", "--headless=true", "--api-version=2", "--log", "--", "-port", "1635" ]
#CMD ["go", "run", "/app", "-port", "1635"]