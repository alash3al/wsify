FROM golang:alpine

RUN apk update && apk add git

# fix error when use docker-compose.yml
RUN go env -w GO111MODULE=off

RUN go get github.com/alash3al/wsify

ENTRYPOINT ["wsify"]

WORKDIR /root/