FROM golang:1.22-alpine As builder

WORKDIR /wsify/

RUN apk update && apk add git upx

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /usr/bin/wsify .

RUN upx -9 /usr/bin/wsify

FROM alpine

WORKDIR /wsify/

COPY --from=builder /usr/bin/wsify /usr/bin/wsify
