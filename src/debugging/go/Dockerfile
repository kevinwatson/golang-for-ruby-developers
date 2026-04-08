# Dockerfile
FROM golang:1.25

WORKDIR /opt/app
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go mod init example.com/employees
COPY employees.go ./
RUN go get ./...

