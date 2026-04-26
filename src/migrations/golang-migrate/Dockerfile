# Dockerfile
FROM golang:1.25
RUN apt update && apt install sqlite3
RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
WORKDIR /opt/app
COPY . .
