FROM node:14 AS BUILDFRONT

WORKDIR /app
COPY ./frontend /app

RUN yarn
RUN yarn build

FROM golang:1.16-alpine AS BUILDBACK

WORKDIR /app
COPY . /app
COPY --from=BUILDFRONT /app/build /app/frontend/build

RUN go mod vendor
RUN go mod download
RUN go build -o server

FROM alpine

RUN apk update && apk add imagemagick && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=BUILDBACK /app/server /app/server

CMD [ "/app/server" ]
