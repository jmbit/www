FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/www-jmbit-de
COPY . .
RUN go get -d -v
RUN go build -o /go/bin/www

FROM scratch
COPY --from=builder /go/bin/www /go/bin/www
ENTRYPOINT ["/go/bin/www"]
