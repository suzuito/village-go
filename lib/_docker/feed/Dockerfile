FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY pkg ./pkg
COPY internal ./internal
COPY cmd ./cmd
RUN go mod download

RUN go build -o /main.exe ./cmd/feed/*.go

FROM alpine:latest

WORKDIR /

COPY --from=build /main.exe /main.exe

ENTRYPOINT [ "/main.exe" ]