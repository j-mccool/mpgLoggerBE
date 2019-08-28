FROM golang:alpine AS builder

RUN apk update && apk add --no-cach git

ARG SOURCE_LOCATION=/
WORKDIR ${SOURCE_LOCATION}

RUN go get -d github.com/gorilla/mux \
    && go get -d go.mongodb.org/mongo-driver/bson \
    && go get -d go.mongodb.org/mongo-driver/bson/primitive \
    && go get -d go.mongodb.org/mongo-driver/mongo \
    && go get -d go.mongodb.org/mongo-driver/mongo/options \
    && go get -d github.com/gorilla/handlers \
    && go get -d golang.org/x/crypto/bcrypt \
    && go get -d github.com/dgrijalva/jwt-go

COPY src /src
COPY main.go .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o mpgBE

##Step 2 - build smaller image
FROM alpine:latest

COPY --from=builder mpgBE .

ENTRYPOINT [ "./mpgBE" ]