FROM golang:1.25

WORKDIR /usr/src/app

RUN go mod init console/main
#RUN go get -u gorm.io/gorm # add these to the go.mod file?
#RUN go get -u gorm.io/driver/sqlite

RUN echo "gorm.io/gorm" >> go.mod
RUN echo "gorm.io/driver/sqlite" >> go.mod


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

