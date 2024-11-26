FROM golang:alpine AS builder
RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community hugo

WORKDIR $GOPATH/src/www-jmbit-de
COPY . .
RUN go get -d -v
RUN hugo --minify
UN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/www

FROM scratch
COPY --from=builder /go/bin/www /go/bin/www
ENTRYPOINT ["/go/bin/www"]
