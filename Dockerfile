FROM golang:1.14-alpine

# the image should only be used for the API side
ARG ARG_USERNAME_DB
ARG ARG_PASSWORD_DB

ENV USERNAME_DB=$ARG_USERNAME_DB
ENV PASSWORD_DB=$ARG_PASSWORD_DB

RUN apk add --update git gcc libc-dev

WORKDIR /go/src/app
COPY monitoring-api .

RUN go get -d -v ./...
RUN go install -v ./...


CMD ["app"]