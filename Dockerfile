FROM node:14 AS BUILDFRONT

WORKDIR /app
COPY ./frontend /app

RUN yarn
RUN yarn build

FROM golang:1.16-alpine AS BUILDBACK

WORKDIR /app
COPY . /app
COPY --from=BUILDFRONT /app/dist /app/frontend/dist

RUN go mod vendor
RUN go mod download
RUN go build -o server

FROM alpine
WORKDIR /app
COPY --from=BUILDBACK /app/server /app/server

CMD [ "/app/server" ]
