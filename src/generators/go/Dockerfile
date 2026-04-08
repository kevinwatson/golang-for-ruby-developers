# Dockerfile
FROM golang:1.25

WORKDIR /opt/app
RUN go mod init example.com/my-app
COPY . .
RUN go get ./...

