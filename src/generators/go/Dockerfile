# Dockerfile
FROM golang:1.25

WORKDIR /opt/app
RUN go mod init example.com/my-app
COPY . .
#RUN go get ./...
RUN go get -u gorm.io/gen@latest
RUN go get -u golang.org/x/tools@latest
RUN go mod tidy

