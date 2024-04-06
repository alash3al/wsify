FROM golang:alpine

RUN apk update && apk add git

RUN go install github.com/alash3al/wsify@latest

ENTRYPOINT ["wsify"]

WORKDIR /root/