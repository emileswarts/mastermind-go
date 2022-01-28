FROM golang:1.17.5-alpine3.15

WORKDIR /app

RUN apk update && apk add build-base bash git
RUN go get github.com/stretchr/testify

CMD ["./scripts/bootstrap.sh"]