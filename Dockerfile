# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go mod download

COPY Dumps ./Dumps
COPY *.go ./


RUN go build -o /main-build

CMD [ "/main-build" ]